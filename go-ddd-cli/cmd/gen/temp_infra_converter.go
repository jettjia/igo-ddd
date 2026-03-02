package gen

const TempInfrastructureConverter = `
package user

import (
	"time"

	"github.com/jinzhu/copier"

	entityUser "xdata/xtext/zsvc-example/domain/entity/user"
	poUser "xdata/xtext/zsvc-example/infra/repository/po/user"
)

// E2PSysMenuAdd entity数据转换成数据库po
func E2PSysMenuAdd(en *entityUser.SysMenu) *poUser.SysMenu {
	var po poUser.SysMenu
	po.CreatedAt = time.Now().UnixMilli()
	if err := copier.Copy(&po, &en); err != nil {
		panic(any(err))
	}

	return &po
}

// E2PSysMenuDel entity数据转换成数据库po
func E2PSysMenuDel(en *entityUser.SysMenu) *poUser.SysMenu {
	var po poUser.SysMenu
	po.DeletedBy = en.DeletedBy
	po.DeletedAt = time.Now().UnixMilli()

	return &po
}

// E2PSysMenuUpdate entity数据转换成数据库po
func E2PSysMenuUpdate(en *entityUser.SysMenu) *poUser.SysMenu {
	var po poUser.SysMenu
	if err := copier.Copy(&po, &en); err != nil {
		panic(any(err))
	}
	po.Ulid = "" // 将使其在 SQL 生成时被忽略, ulid是主键

	po.UpdatedAt = time.Now().UnixMilli()
	return &po
}

// P2ESysMenu 数据库po转换成entity
func P2ESysMenu(po *poUser.SysMenu) *entityUser.SysMenu {
	var entity entityUser.SysMenu
	if err := copier.Copy(&entity, &po); err != nil {
		panic(any(err))
	}

	return &entity
}

func P2ESysMenus(pos []*poUser.SysMenu) []*entityUser.SysMenu {
	ens := make([]*entityUser.SysMenu, 0)
	if len(pos) == 0 {
		return ens
	}

	for _, val := range pos {
		cfg := P2ESysMenu(val)
		ens = append(ens, cfg)
	}

	return ens
}

`
