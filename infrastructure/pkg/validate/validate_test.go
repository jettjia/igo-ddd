package validate

import (
	"fmt"
	"testing"
)

func Test_CheckSpecialCharacters(t *testing.T) {
	type User struct {
		UserName string `json:"user_name"` //通过reg_error_info标签记录
		//reg_error_info也可以是标记错误的唯一标识，通过传入的local_language 从库中或者缓存中找到对应国家的错误提示信息
		Password string `json:"password" validate:"specialChara" err_info:"密码至少6个字符"`
	}

	var list []User

	user1 := User{UserName: "zzzzzz", Password: "xxxx."}
	user2 := User{UserName: "zzzzzz", Password: "xxxx@"}
	user3 := User{UserName: "zzzzzz", Password: "xxxx-"}
	user4 := User{UserName: "zzzzzz", Password: "xxxx_"}
	user5 := User{UserName: "zzzzzz", Password: "xxxx.@"}
	user6 := User{UserName: "zzzzzz", Password: "xxxx.中"}
	user7 := User{UserName: "zzzzzz", Password: "xxxx*"}
	user8 := User{UserName: "zzzzzz", Password: "xxxx$"}
	user9 := User{UserName: "zzzzzz", Password: "capp"}
	user10 := User{UserName: "zzzzzz", Password: "capp_01"}

	list = append(list, user1, user2, user3, user4, user5, user6, user7, user8, user9, user10)

	for _, v := range list {
		err := Validate(v) //校验

		if err != nil {
			fmt.Println(fmt.Sprintf("当前校验单字符串是：【%s】, 错误是：%s", v.Password, err.Error()))
		} else {
			fmt.Println(fmt.Sprintf("当前校验单字符串是：【%s】, 成功", v.Password))
		}
	}
}
