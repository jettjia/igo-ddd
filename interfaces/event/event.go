package event

import (
	"github.com/jett/gin-ddd/cmd"
	nsq "github.com/jett/gin-ddd/interfaces/event/subscribe"
)

// InitEvent 消息事件
func InitEvent(app *cmd.App) {
	nsq.UserSubDemo(app)
}
