package nsq

import "strconv"

// UserNsqSubDemo 处理接受到的消息的逻辑
func UserNsqSubDemo(id uint64) string {
	return "domain get id is :" + strconv.Itoa(int(id))
}
