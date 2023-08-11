package redis

import (
	"log"
	"net"

	"bitbucket.org/non-pn/mini-redis-go/internal/constant"
	"bitbucket.org/non-pn/mini-redis-go/internal/network"
	"bitbucket.org/non-pn/mini-redis-go/internal/utils"
)

type RedisClient struct {
	Network    string
	Host       string
	Connection *net.Conn
}

func NewClient(host string) *RedisClient {
	rediscli := &RedisClient{
		Network: constant.PROTOCOL,
		Host:    host,
	}

	return rediscli
}

func (cli *RedisClient) Connect() error {
	conn, err := net.Dial(cli.Network, cli.Host)
	if err != nil {
		return err
	}

	cli.Connection = &conn

	return nil
}

func (cli *RedisClient) Close() error {
	err := (*cli.Connection).Close()
	if err != nil {
		return err
	}

	return nil
}

func (cli *RedisClient) Send(data []byte) ([]byte, error) {
	err := network.WriteWithCRLF(cli.Connection, data)
	if err != nil {
		return nil, err
	}

	resp, err := network.ReadUntilCRLF(cli.Connection)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (cli *RedisClient) SendGetCmd(k string) (*RedisResponsePayload, error) {
	payload := &RedisRequestPayload{
		Cmd:   GetCmd,
		Key:   k,
		Value: []byte{},
	}
	rawpayload, err := payload.ToRaw()
	if err != nil {
		return nil, err
	}

	rawresp, err := cli.Send(*rawpayload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rawredisresp := RedisRawResponsePayload(rawresp)
	redisresp, err := rawredisresp.TransformPayload()
	if err != nil {
		return nil, err
	}

	return redisresp, nil
}

func (cli *RedisClient) SendSetCmd(k string, v utils.TypeLengthValue) (*RedisResponsePayload, error) {
	payload := &RedisRequestPayload{
		Cmd:   SetCmd,
		Key:   k,
		Value: v,
	}
	rawpayload, err := payload.ToRaw()
	if err != nil {
		return nil, err
	}

	rawresp, err := cli.Send(*rawpayload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rawredisresp := RedisRawResponsePayload(rawresp)
	redisresp, err := rawredisresp.TransformPayload()
	if err != nil {
		return nil, err
	}

	return redisresp, nil
}
