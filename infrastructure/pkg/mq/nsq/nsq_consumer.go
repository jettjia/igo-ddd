package nsq

import (
	"github.com/nsqio/go-nsq"

	"github.com/jettjia/gin-ddd/global"
)

const (
	ChannelNsq string = "gin-ddd"
)

type Consumer interface {
	SubscribeMsg(topic string, cmd func([]byte) error) error
}

var _ Consumer = (*consumer)(nil)

type consumer struct {
}

func NewConsumerClient() *consumer {
	return &consumer{}
}

type mqHandler struct {
	cmd func([]byte) (err error)
}

// HandleMessage 实现消息处理函数
func (mq *mqHandler) HandleMessage(m *nsq.Message) (err error) {
	if len(m.Body) == 0 {
		return nil
	}
	// 将消息体传给实际处理的方法
	err = mq.cmd(m.Body)
	return
}

func (s *consumer) SubscribeMsg(topic string, cmd func([]byte) error) error {
	client := NewNsqClient()
	cc, err := client.NsqConsumer(topic, ChannelNsq)
	if err != nil {
		global.GLog.Errorf("ERROR: NewConsumer:%v\n", err)
	}

	handler := &mqHandler{cmd: cmd}
	cc.AddHandler(handler)

	address := client.GetNsqConsumerLink()
	if err := cc.ConnectToNSQD(address); err != nil { // 直接连NSQD
		//if err := cc.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询
		return err
	}

	return nil
}
