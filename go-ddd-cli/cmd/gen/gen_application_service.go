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

func GenApplicationService(req GenReq) {
	path := req.AServiceDir + "/" + req.TableName + "_svc.go"

	context2 := gstr.ReplaceByMap(TempApplicationService, g.MapStrStr{
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"assembler/user":           "assembler/" + viper.Get("server.server_name").(string),
		"dto/user":                 "dto/" + viper.Get("server.server_name").(string),
		"srv/user":                 "srv/" + viper.Get("server.server_name").(string),
		"aggregate/user":           "aggregate/" + viper.Get("server.server_name").(string),
		"SysMenu":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"sysMenu":                  GetJsonTagFromCase(req.TableName, "CamelLower"),
		"package user":             "package " + viper.Get("server.server_name").(string),
		"entityUser":               "entity" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
		"dtoUser":                  "dto" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
		"assUser":                  "ass" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
		"aggUser":                  "agg" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
		"srvUser":                  "srv" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context2)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
