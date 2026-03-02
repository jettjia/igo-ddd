package apierror

import "github.com/jettjia/igo-pkg/pkg/xerror"

var (
	i18nInfo = make(map[int]string)
)

var (
	// 推算规则，比如定义的服务端口是：34100
	// 那么错误码就是：400 100 1001 这种规则
	SystemCode = 1000 // 系统错误码

	// 通用错误码
	BadRequestErr     = 400000000 + SystemCode
	UnauthorizedErr   = 401000000 + SystemCode
	ForbiddenErr      = 403000000 + SystemCode
	NotFoundErr       = 404000000 + SystemCode
	ConflictErr       = 409000000 + SystemCode
	InternalServerErr = 500000000 + SystemCode

	// 用户模块错误
	UserPwdOrNicknameErr = UnauthorizedErr + 001 // 用户账户或密码错误

	UserCode            = 100                          // 用户模块错误码
	UserNotFoundErr     = BadRequestErr + UserCode + 1 // 400100201
	TooManyRequestsErr  = BadRequestErr + UserCode + 9 // 400100209
	UserNameConflictErr = ConflictErr + UserCode + 2   // 409100102

	// 菜单模块错误
	MenuCode            = 200                          // 用户模块错误码
	MenuNotFoundErr     = BadRequestErr + MenuCode + 1 // 400100201
	MenuNameConflictErr = ConflictErr + MenuCode + 2   // 409100202
)

// 通用异常
func init() {
	// 通用错误码
	i18nInfo[BadRequestErr] = "BadRequestErr"
	i18nInfo[UnauthorizedErr] = "UnauthorizedErr"
	i18nInfo[ForbiddenErr] = "ForbiddenErr"
	i18nInfo[NotFoundErr] = "NotFoundErr"
	i18nInfo[ConflictErr] = "ConflictErr"
	i18nInfo[InternalServerErr] = "InternalServerErr"

	// 官网模块
	i18nInfo[TooManyRequestsErr] = "TooManyRequestsErr"

	// 用户模块
	i18nInfo[UserPwdOrNicknameErr] = "UserPwdOrNicknameErr"

	// 菜单模块错误
	i18nInfo[MenuNotFoundErr] = "MenuNotFoundErr"
	i18nInfo[MenuNameConflictErr] = "MenuNameConflictErr"

	xerror.LoadTranslation(i18nInfo)
}
