package user

import (
	"time"

	"github.com/jinzhu/copier"

	entityUser "github.com/jettjia/ddddemo/domain/entity/user"
	poUser "github.com/jettjia/ddddemo/infra/repository/po/user"
)

// E2PSysLogAdd entity数据转换成数据库po
func E2PSysLogAdd(en *entityUser.SysLog) *poUser.SysLog {
	var po poUser.SysLog
	po.CreatedAt = time.Now().UnixNano()
	if err := copier.Copy(&po, &en); err != nil {
		panic(any(err))
	}

	return &po
}

// E2PSysLogDel entity数据转换成数据库po
func E2PSysLogDel(en *entityUser.SysLog) *poUser.SysLog {
	var po poUser.SysLog
	po.DeletedBy = en.DeletedBy
	po.DeletedAt = time.Now().UnixMilli()

	return &po
}

// E2PSysLogUpdate entity数据转换成数据库po
func E2PSysLogUpdate(en *entityUser.SysLog) *poUser.SysLog {
	var po poUser.SysLog
	if err := copier.Copy(&po, &en); err != nil {
		panic(any(err))
	}

	po.UpdatedAt = time.Now().UnixNano()
	return &po
}

// P2ESysLog 数据库po转换成entity
func P2ESysLog(po *poUser.SysLog) *entityUser.SysLog {
	var entity entityUser.SysLog
	if err := copier.Copy(&entity, &po); err != nil {
		panic(any(err))
	}

	return &entity
}

func P2ESysLogs(pos []*poUser.SysLog) []*entityUser.SysLog {
	ens := make([]*entityUser.SysLog, 0)
	if len(pos) == 0 {
		return ens
	}

	for _, val := range pos {
		cfg := P2ESysLog(val)
		ens = append(ens, cfg)
	}

	return ens
}
