package event

import (
	"github.com/jettjia/go-ddd/cmd"
	nsq "github.com/jettjia/go-ddd/interfaces/event/subscribe"
)

// InitEvent 消息事件
func InitEvent(app *cmd.App) {
	nsq.UserSubDemo(app)
}
