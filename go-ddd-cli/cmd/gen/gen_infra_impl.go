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

func GenInfraImpl(req GenReq) {
	context := gstr.ReplaceByMap(TempInfraImpl, g.MapStrStr{
		// import包
		"package user":             "package " + viper.Get("server.server_name").(string),
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),

		"entity/user":      "entity/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"irepository/user": "irepository/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"converter/user":   "converter/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"po/user":          "po/" + viper.Get("server.server_name").(string),
		"sysMenuEn":        GetJsonTagFromCase(req.TableName, "CamelLower") + "En",

		"entityUser":      "entity" + GetJsonTagFromCase(req.SrvName, "Camel"),
		"irepositoryUser": "irepository" + GetJsonTagFromCase(req.TableName, "Camel"), // 表名
		"converterUser":   "converter" + GetJsonTagFromCase(req.TableName, "Camel"),   // 表名
		"poUser":          "po" + GetJsonTagFromCase(req.TableName, "Camel"),          // 表名

		// 参数
		"SysMenu":   GetJsonTagFromCase(req.TableName, "Camel"),
		"sys_menu":  GetJsonTagFromCase(req.TableName, "Snake"), // 表名
		"sysMenuPo": GetJsonTagFromCase(req.TableName, "CamelLower") + "Po",
	})

	path := req.RepositoryimplDir + "/" + req.TableName + "_impl.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
