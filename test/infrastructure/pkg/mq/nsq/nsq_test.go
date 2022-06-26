package nsq

import (
	"fmt"
	"testing"

	_ "github.com/jett/gin-ddd/boot"
	"github.com/jett/gin-ddd/infrastructure/pkg/mq/nsq"
)

func TestNsq_NewNsqProducer(t *testing.T) {
	client := nsq.NewNsqClient()
	cp, err := client.NsqProducer()
	if err != nil {
		t.Fatal(err)
	}
	err = cp.Ping()
	if nil != err {
		// 关闭生产者
		cp.Stop()
		cp = nil
		t.Fatal(err)
	}

	fmt.Println("ping nsq success")
}
