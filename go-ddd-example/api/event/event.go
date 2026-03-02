package event

import (
	subscribeUser "github.com/jettjia/ddddemo/api/event/subscribe/user"
	"github.com/jettjia/ddddemo/config"
)

// InitEvent 消息事件
func InitEvent() {
	conf := config.NewConfig()
	if !conf.Server.EnableEvent {
		return
	}

	sub := subscribeUser.NewSubscribe()
	sub.SubscribeCreateUser()
}
