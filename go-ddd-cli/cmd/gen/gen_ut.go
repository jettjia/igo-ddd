package gen

import (
	"strings"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/spf13/viper"
)

func GenUt(req GenReq) {
	return // 这里应该是svc的单测
	context := gstr.ReplaceByMap(TempUt, g.MapStrStr{
		"GasPc":            viper.Get("server.go_module").(string),
		"entity/user":      "entity/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"irepository/user": "irepository/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"SysMenu":          GetJsonTagFromCase(req.TableName, "Camel"),
		"sysMenu":          GetJsonTagFromCase(req.TableName, "CamelLower"), // 表名
		"MenuId":           GetJsonTagFromCase(req.TableKey, "Camel"),       // 主键
	})

	path := req.AggregateDir + "/" + req.TableName + "_agg_test.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
