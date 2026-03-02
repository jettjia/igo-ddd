package gen

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func GenDomainAgg(req GenReq) {
	context := gstr.ReplaceByMap(TempDomainAgg, g.MapStrStr{
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"entity/user":              "entity/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"repo/user":                "repo/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"irepository/user":         "irepository/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"repositoryimpl/user":      "repositoryimpl/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"SysMenu":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"sys_menu":                 GetJsonTagFromCase(req.TableName, "Snake"),      // 表名
		"sysMenu":                  GetJsonTagFromCase(req.TableName, "CamelLower"), // 表名
		"package user":             "package " + viper.Get("server.server_name").(string),
		"entityUser":               "entity" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
		"repoUser":                 "repo" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
	})

	path := req.AggregateDir + "/" + req.TableName + "_agg.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
