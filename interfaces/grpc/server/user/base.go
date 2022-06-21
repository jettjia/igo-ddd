package user

import (
	service "github.com/jett/gin-ddd/application/service/user"
)

type UserServer struct {
	UserSrv *service.UserService // 注入 application/service
}
