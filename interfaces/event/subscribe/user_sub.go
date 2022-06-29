package nsq

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"

	dto "github.com/jettjia/go-ddd/application/dto/user"
	"github.com/jettjia/go-ddd/cmd"
	"github.com/jettjia/go-ddd/global"
	"github.com/jettjia/go-ddd/infrastructure/pkg/mq/nsq"
)

// UserSubDemo subscribe mq demo
func UserSubDemo(app *cmd.App) {
	// todo, nsq error
	client := nsq.NewConsumerClient()
	err := client.SubscribeMsg("test_topic", func(bytes []byte) error {
		// 输出 nsq 中的消息
		jsonStr := string(bytes)
		gResult := gjson.Get(jsonStr, "id")

		// 在这里模拟你要处理的逻辑，处理你的业务，调用 application/service， application/service --> domain(event->subscribe)
		userInfo, _ := app.UserSrv.GetSimpleUserInfo(context.Background(), &dto.SimpleUserInfoReq{
			Id: uint64(gResult.Int()),
		})
		fmt.Println("从 nsq subscribe获取到的用户信息是 userInfo: ", userInfo)
		return nil
	})

	if err != nil {
		global.GLog.Errorln(err)
	}
}
