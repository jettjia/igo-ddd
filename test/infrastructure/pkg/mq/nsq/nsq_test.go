package nsq

import (
	"fmt"
	"testing"

	_ "github.com/jett/gin-ddd/boot"
	"github.com/jett/gin-ddd/global"
)

func TestNsq_NewNsqProducer(t *testing.T) {
	err := global.GNsqProducer.Ping()
	if nil != err {
		// 关闭生产者
		global.GNsqProducer.Stop()
		global.GNsqProducer = nil
		t.Fatal(err)
	}

	fmt.Println("ping nsq success")
}
