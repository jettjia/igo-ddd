package converter

import (
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/plugin/soft_delete"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/util"
	po "github.com/jettjia/go-ddd-demo/infrastructure/repository/po/sys"
)

// E2PSysMenuAdd entity数据转换成数据库po
func E2PSysMenuAdd(en *entity.SysMenu) *po.SysMenu {
	var sysMenu po.SysMenu
	sysMenu.CreatedAt = time.Now().UnixNano()
	sysMenu.Uuid = util.Ulid()
	if err := copier.Copy(&sysMenu, &en); err != nil {
		panic(any(err))
	}

	return &sysMenu
}

// E2PSysMenuDel entity数据转换成数据库po
func E2PSysMenuDel(en *entity.SysMenu) *po.SysMenu {
	var sysMenu po.SysMenu
	sysMenu.DeletedBy = en.DeletedBy
	sysMenu.DeletedAt = soft_delete.DeletedAt(time.Now().UnixNano())

	return &sysMenu
}

// E2PSysMenuUpdate entity数据转换成数据库po
func E2PSysMenuUpdate(en *entity.SysMenu) *po.SysMenu {
	var sysMenu po.SysMenu
	if err := copier.Copy(&sysMenu, &en); err != nil {
		panic(any(err))
	}

	sysMenu.UpdatedAt = time.Now().UnixNano()
	return &sysMenu
}

// P2ESysMenu 数据库po转换成entity
func P2ESysMenu(po *po.SysMenu) *entity.SysMenu {
	var sysMenu entity.SysMenu
	if err := copier.Copy(&sysMenu, &po); err != nil {
		panic(any(err))
	}

	return &sysMenu
}

func P2ESysMenus(pos []*po.SysMenu) []*entity.SysMenu {
	ens := make([]*entity.SysMenu, 0)
	if len(pos) == 0 {
		return ens
	}

	for _, val := range pos {
		cfg := P2ESysMenu(val)
		ens = append(ens, cfg)
	}

	return ens
}
