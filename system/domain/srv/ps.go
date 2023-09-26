package srv

import "github.com/google/wire"

var SvcProviderSet = wire.NewSet(NewSysMenuSvc, NewSysLogSvc)
