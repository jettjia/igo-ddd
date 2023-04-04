package converter

import (
	"time"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/util"
	po "github.com/jettjia/go-ddd-demo/infrastructure/repository/po/sys"
	"github.com/jinzhu/copier"
)

// E2PSysLogAdd entity数据转换成数据库po
func E2PSysLogAdd(en *entity.SysLog) *po.SysLog {
	var sysLog po.SysLog
	if err := copier.Copy(&sysLog, &en); err != nil {
		panic(any(err))
	}
	sysLog.CreatedAt = time.Now().UnixMilli()
	sysLog.Uuid = util.Ulid()

	return &sysLog
}

// E2PSysLogDel entity数据转换成数据库po
func E2PSysLogDel(en *entity.SysLog) *po.SysLog {
	var sysLog po.SysLog

	return &sysLog
}

// E2PSysLogUpdate entity数据转换成数据库po
func E2PSysLogUpdate(en *entity.SysLog) *po.SysLog {
	var sysLog po.SysLog
	if err := copier.Copy(&sysLog, &en); err != nil {
		panic(any(err))
	}

	return &sysLog
}

// P2ESysLog 数据库po转换成entity
func P2ESysLog(po *po.SysLog) *entity.SysLog {
	var sysLog entity.SysLog
	if err := copier.Copy(&sysLog, &po); err != nil {
		panic(any(err))
	}

	return &sysLog
}

func P2ESysLogs(pos []*po.SysLog) []*entity.SysLog {
	ens := make([]*entity.SysLog, 0)
	if len(pos) == 0 {
		return ens
	}

	for _, val := range pos {
		cfg := P2ESysLog(val)
		ens = append(ens, cfg)
	}

	return ens
}
