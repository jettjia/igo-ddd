package user

import (
	"github.com/jettjia/igo-pkg/pkg/xsql/builder"
)

// 请求对象
type (
	// CreateSysUserReq 创建SysUser 请求对象
	CreateSysUserReq struct {
		CreatedBy  string `json:"created_by"`                      // 创建者
		MemberCode string `json:"member_code" validate:"required"` // 会员号
		Phone      string `json:"phone"`                           // 手机号码
		NickName   string `json:"nick_name"`                       // 昵称
		TrueName   string `json:"true_name"`                       // 真实姓名，可重复
		LevelId    string `json:"level_id"`                        // 会员等级id
		Password   string `json:"password"`                        // 密码
		AdminLevel uint   `json:"admin_level"`                     // 1是admin超管

		Exp CreateSysUserReqExp `json:"exp"`
	}

	CreateSysUserReqExp struct {
		Addr     string `json:"addr"`      // 住址
		AddrCode string `json:"addr_code"` // 住址编号
	}

	// DelSysUsersReq 删除 请求对象
	DelSysUsersReq struct {
		Ulid      string `validate:"required" uri:"ulid" json:"ulid"` // ulid
		DeletedBy string `json:"deleted_by"`                          // 删除者

	}

	// UpdateSysUserReq 修改SysUser 请求对象
	UpdateSysUserReq struct {
		Ulid       string `validate:"required" uri:"ulid" json:"ulid"` // ulid
		UpdatedBy  string `json:"updated_by"`                          // 修改者
		MemberCode string `json:"member_code"`                         // 会员号
		Phone      string `json:"phone"`                               // 手机号码
		NickName   string `json:"nick_name"`                           // 昵称
		Unionid    string `json:"unionid"`                             // 微信unionid
		LevelId    string `json:"level_id"`                            // 会员等级id
	}

	// FindSysUserByIdReq 查询 请求对象
	FindSysUserByIdReq struct {
		Ulid string `validate:"required" uri:"ulid" json:"ulid"` // ulid
	}

	// FindSysUserByQueryReq 查询 请求对象
	FindSysUserByQueryReq struct {
		Query []*builder.Query `json:"query"`
	}

	// FindSysUserAllReq 查询 请求对象
	FindSysUserAllReq struct {
		Query []*builder.Query `json:"query"`
	}

	// FindSysUserPageReq 分页查询 请求对象
	FindSysUserPageReq struct {
		Query    []*builder.Query  `json:"query"`
		PageData *builder.PageData `json:"page_data"`
		SortData *builder.SortData `json:"sort_data"`
	}

	// LoginReq 登录
	LoginReq struct {
		NickName string `validate:"required" json:"name"`     // 昵称
		Password string `validate:"required" json:"password"` // 密码
		LoginIp  string `json:"login_ip"`                     // 登录Ip
	}

	// LoginReq 登出
	LogoutReq struct {
		Token string `json:"token"` // 登录的token
	}
)

// 输出对象
type (

	// LoginRsp 登录rsp
	LoginRsp struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		UserExp      string `json:"user_exp"`
	}

	// CreateSysUserRsp 创建SysUser 返回对象
	CreateSysUserRsp struct {
		Ulid string `json:"ulid"` // ulid

	}

	// FindSysUserPageRsp 列表查询 返回对象
	FindSysUserPageRsp struct {
		Entries  []*FindSysUserRsp `json:"entries"`
		PageData *builder.PageData `json:"page_data"`
	}

	// FindSysUserRsp 查询SysUser 返回对象
	FindSysUserRsp struct {
		Ulid       string `json:"ulid"`        // ulid
		CreatedAt  int64  `json:"created_at"`  // 创建时间
		UpdatedAt  int64  `json:"updated_at"`  // 修改时间
		CreatedBy  string `json:"created_by"`  // 创建者
		UpdatedBy  string `json:"updated_by"`  // 修改者
		MemberCode string `json:"member_code"` // 会员号
		Phone      string `json:"phone"`       // 手机号码
		NickName   string `json:"nick_name"`   // 昵称
		Unionid    string `json:"unionid"`     // 微信unionid
		LevelId    string `json:"level_id"`    // 会员等级id

		Exp CreateSysUserRspExp `json:"exp"`
	}

	CreateSysUserRspExp struct {
		Addr     string `json:"addr"`      // 住址
		AddrCode string `json:"addr_code"` // 住址编号
	}
)
