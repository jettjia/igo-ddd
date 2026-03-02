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

func GenDomainSvc(req GenReq) {

	tempStr := TempDomainSvc
	// 判断是不是有删除字段
	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" {
			tempStr = TempDomainSvc2
			break
		}
	}

	context := gstr.ReplaceByMap(tempStr, g.MapStrStr{
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

	path := req.ServiceDir + "/" + req.TableName + "_svc.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
