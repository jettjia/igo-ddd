package nsq

import (
	"github.com/nsqio/go-nsq"

	"github.com/jett/gin-ddd/global"
)

// NewNsqProducer return *nsq.Producer
func NewNsqProducer() *nsq.Producer {
	addr := getNsqProducerLink()
	nsq, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	return nsq
}

// GetNsqProducerLink return nsq link
func getNsqProducerLink() string {
	return global.Gconfig.Nsq.NsqProducerHost + ":" + global.Gconfig.Nsq.NsqProducerPort
}

func NewNsqConsumer(topic string, ch string) (*nsq.Consumer, error) {
	config := nsq.NewConfig()
	return nsq.NewConsumer(topic, ch, config)
}

func NewNsqConn(addr string, config *nsq.Config, delegate nsq.ConnDelegate) *nsq.Conn {
	return nsq.NewConn(addr, config, delegate)
}
