package gen

const TempApplicationDto = `
package user

import (
	"github.com/jettjia/igo-pkg/pkg/xsql/builder"
)

// 请求对象
type (
	// CreateSysMenuReq 创建SysMenu 请求对象
	CreateSysMenuReq struct {
		{{createDto}}
	}

	// DelSysMenusReq 删除 请求对象
	DelSysMenusReq struct {
		{{delDto}}
	}

	// UpdateSysMenuReq 修改SysMenu 请求对象
	UpdateSysMenuReq struct {
		{{updateDto}}
	}

	// FindSysMenuByIdReq 查询 请求对象
	FindSysMenuByIdReq struct {
		{{findByIdDto}}
	}

	// FindSysMenuByQueryReq 查询 请求对象
	FindSysMenuByQueryReq struct {
		{{findQueryDto}}
	}

	// FindSysMenuAllReq 查询 请求对象
	FindSysMenuAllReq struct {
		{{findAllDto}}
	}

	// FindSysMenuPageReq 分页查询 请求对象
	FindSysMenuPageReq struct {
		{{findPageDto}}
	}
)

// 输出对象
type (
	// CreateSysMenuRsp 创建SysMenu 返回对象
	CreateSysMenuRsp struct {
		{{createRsp}}
	}

	// FindSysMenuPageRsp 列表查询 返回对象
	FindSysMenuPageRsp struct {
		{{pageRsp}}
	}

	// FindSysMenuRsp 查询SysMenu 返回对象
	FindSysMenuRsp struct {
		{{findRsp}}
	}
)

`
