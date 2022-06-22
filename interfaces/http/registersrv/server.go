package registersrv

import service "github.com/jett/gin-ddd/application/service/user"

type Registersrv struct {
	UserSrv *service.UserService // 注入 application/service
}
