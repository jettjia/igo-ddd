package nsq

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jett/gin-ddd/cmd"
	"github.com/jett/gin-ddd/global"
	"github.com/jett/gin-ddd/infrastructure/pkg/mq/nsq"
)

// UserSubDemo subscribe mq demo
func UserSubDemo(app *cmd.App) {
	// todo, nsq error
	client := nsq.NewConsumerClient()
	err := client.SubscribeMsg("test_topic", func(bytes []byte) error {
		// 输出 nsq 中的消息
		fmt.Println("00000000000000000")
		fmt.Println(string(bytes))

		return nil
	})

	if err != nil {
		global.GLog.Errorln(err)
	}

	c := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c
}
