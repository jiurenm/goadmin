package thrift

import (
	"admin/api/keys"
	"admin/pkg/conf"
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
)

type Thrift struct {
	client *keys.GetKeywordsClient
}

func New(yaml *conf.Yaml) (*Thrift, error) {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(yaml.Thrift.Host, yaml.Thrift.Port))
	if err != nil {
		return nil, err
	}
	useTransport, err := transportFactory.GetTransport(transport)
	client := keys.NewGetKeywordsClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		return nil, err
	}
	return &Thrift{client: client}, nil
}

func (t *Thrift) GetKeyWord(sentence string) string {
	res, err := t.client.Get(context.Background(), sentence)
	if err != nil {
		return "其他"
	}
	if res == "" {
		return "其他"
	}
	return res
}
