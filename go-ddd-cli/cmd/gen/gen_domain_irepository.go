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

func GenDomainIrepository(req GenReq) {
	context := gstr.ReplaceByMap(TempDomainIrepository, g.MapStrStr{
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"entity/user":              "entity/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"entityUser":               "entity" + GetJsonTagFromCase(req.SrvName, "Camel"),
		"SysMenu":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"sys_menu":                 GetJsonTagFromCase(req.TableName, "Snake"),             // 表名
		"sysMenuEn":                GetJsonTagFromCase(req.TableName, "CamelLower") + "En", // 表名
		"uint64":                   GetPriIdType(req),                                      // 主键id
		"package user":             "package " + viper.Get("server.server_name").(string),
	})

	path := req.IrepositoryDir + "/" + "i_" + req.TableName + "_repo.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
