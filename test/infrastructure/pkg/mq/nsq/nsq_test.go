package nsq

import (
	"fmt"
	"testing"

	"github.com/jettjia/gin-ddd/infrastructure/pkg/mq/nsq"
	_ "github.com/jettjia/gin-ddd/test"
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
