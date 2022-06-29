package event

import (
	"github.com/jettjia/gin-ddd/cmd"
	nsq "github.com/jettjia/gin-ddd/interfaces/event/subscribe"
)

// InitEvent 消息事件
func InitEvent(app *cmd.App) {
	nsq.UserSubDemo(app)
}
