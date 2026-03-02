package util

import (
	"fmt"
	"jettjia/igo-ddd-cli/util"
	"testing"
)

func Test_GetOutboundIP(t *testing.T) {
	str := util.GetOutboundIP()
	fmt.Println("IP地址是：", str)
}
