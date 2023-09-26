package aggregate

import "github.com/google/wire"

var AggProviderSet = wire.NewSet(NewSysMenuAgg)
