package assembler

import "github.com/google/wire"

var AsseProviderSet = wire.NewSet(NewSysMenuReq, NewSysMenuRsp)
