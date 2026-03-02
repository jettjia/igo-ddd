package user

import (
	"time"

	"github.com/jinzhu/copier"

	entityUser "github.com/jettjia/ddddemo/domain/entity/user"
	poUser "github.com/jettjia/ddddemo/infra/repository/po/user"
)

// E2PSysUserAdd entity数据转换成数据库po
func E2PSysUserAdd(en *entityUser.SysUser) *poUser.SysUser {
	var po poUser.SysUser
	po.CreatedAt = time.Now().UnixNano()
	if err := copier.Copy(&po, &en); err != nil {
		panic(any(err))
	}

	return &po
}

// E2PSysUserDel entity数据转换成数据库po
func E2PSysUserDel(en *entityUser.SysUser) *poUser.SysUser {
	var po poUser.SysUser
	po.DeletedBy = en.DeletedBy
	po.DeletedAt = time.Now().UnixMilli()

	return &po
}

// E2PSysUserUpdate entity数据转换成数据库po
func E2PSysUserUpdate(en *entityUser.SysUser) *poUser.SysUser {
	var po poUser.SysUser
	if err := copier.Copy(&po, &en); err != nil {
		panic(any(err))
	}

	po.UpdatedAt = time.Now().UnixNano()
	return &po
}

// P2ESysUser 数据库po转换成entity
func P2ESysUser(po *poUser.SysUser) *entityUser.SysUser {
	var entity entityUser.SysUser
	if err := copier.Copy(&entity, &po); err != nil {
		panic(any(err))
	}

	return &entity
}

func P2ESysUsers(pos []*poUser.SysUser) []*entityUser.SysUser {
	ens := make([]*entityUser.SysUser, 0)
	if len(pos) == 0 {
		return ens
	}

	for _, val := range pos {
		cfg := P2ESysUser(val)
		ens = append(ens, cfg)
	}

	return ens
}
