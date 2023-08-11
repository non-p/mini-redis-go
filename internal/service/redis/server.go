package redis

import (
	"io"
	"log"
	"net"
	"time"

	"bitbucket.org/non-pn/mini-redis-go/internal/constant"
	"bitbucket.org/non-pn/mini-redis-go/internal/db"
	"bitbucket.org/non-pn/mini-redis-go/internal/network"
)

type RedisServer struct {
	Port        string
	Listener    net.Listener
	Connections []*net.Conn
	DB          *db.KVStore[[]byte]
}

func NewServer(port string) (*RedisServer, error) {
	kvstore := db.InitKVStore[[]byte](nil)
	l, err := net.Listen(constant.PROTOCOL, port)
	if err != nil {
		return nil, err
	}

	return &RedisServer{
		Port:        port,
		Listener:    l,
		Connections: make([]*net.Conn, 0, 5),
		DB:          kvstore,
	}, nil
}

func (serv *RedisServer) Start() error {
	for {
		c, err := serv.Listener.Accept()
		if err != nil {
			return err
		}

		log.Println("New connection from:", c.RemoteAddr())
		go serv.HandleConnection(&c)
	}
}

func (serv *RedisServer) Stop() error {
	for _, conn := range serv.Connections {
		err := (*conn).Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (serv *RedisServer) HandleConnection(conn *net.Conn) error {
	serv.Connections = append(serv.Connections, conn)

	for {
		data, err := network.ReadUntilCRLF(conn)
		if err != nil {
			if err == io.EOF {
				log.Println("Client disconnected")
			}
			return err
		}

		reqctx := &network.RequestContext{
			Now:  time.Now(),
			Data: data,
			Conn: conn,
		}
		err = serv.HandleRequest(reqctx)
		if err != nil {
			return err
		}
	}
}

func (serv *RedisServer) HandleRequest(ctx *network.RequestContext) error {
	var resp RedisResponsePayload
	data := RedisRawRequestPayload(ctx.Data)
	payload, err := data.TransformPayload()
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Receive payload:", payload)

	switch payload.Cmd {
	case GetCmd:
		resp = RedisResponsePayload{
			RespType: GetSuccess,
			RespBody: serv.Get(payload.Key),
		}
		raw, err := resp.ToRaw()
		if err != nil {
			log.Println(err)
			return err
		}
		err = ctx.Response(*raw)
		if err != nil {
			log.Println(err)
			return err
		}
		break
	case SetCmd:
		serv.Set(payload.Key, []byte(payload.Value))
		resp = RedisResponsePayload{
			RespType: SetSuccess,
			RespBody: []byte("OK"),
		}
		raw, err := resp.ToRaw()
		if err != nil {
			log.Println(err)
			return err
		}
		err = ctx.Response(*raw)
		if err != nil {
			log.Println(err)
			return err
		}
		break
	}

	err = serv.DB.CacheStorage()
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (serv *RedisServer) Get(k string) []byte {
	log.Printf("Get from KVStore with k[%v]", k)
	return serv.DB.Get(k)
}

func (serv *RedisServer) Set(k string, v []byte) {
	log.Printf("Set to KVStore with k[%v] v[%v]\n", k, v)
	serv.DB.Set(k, v)
}
