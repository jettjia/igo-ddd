package ievent

import (
	"github.com/jettjia/ddddemo/config"

	"github.com/jettjia/igo-pkg/pkg/event"
)

func NewEventCli() event.MQClient {
	conf := config.NewConfig()

	return event.NewMQClient(conf.Mq.MqProducerHost, conf.Mq.MqProducerPort, conf.Mq.MqSubscribeHost, conf.Mq.MqSubscribePort, conf.Mq.MqType)
}
