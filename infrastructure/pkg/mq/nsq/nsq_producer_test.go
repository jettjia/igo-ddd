package nsq

import (
	"testing"

	_ "github.com/jettjia/gin-ddd/boot"
)

func TestNsq_Publish(t *testing.T) {
	msg := make(map[string]string)
	msg["name"] = "jettjiajia"

	client := NewPublishClient()
	err := client.PublishMsg("test_topic", msg)
	if err != nil {
		t.Fatal(err)
	}
}
