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

func GenApplicationAssembler(req GenReq) {
	genAssDto(req)
}

func genAssDto(req GenReq) {
	path := req.AssemblerDir + "/" + req.TableName + "_assembler.go"

	context2 := gstr.ReplaceByMap(TempAssemblerTdo, g.MapStrStr{
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"SysMenu":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"sysMenu":                  GetJsonTagFromCase(req.TableName, "CamelLower"),
		"dto/user":                 "dto/" + viper.Get("server.server_name").(string),
		"entity/user":              "entity/" + viper.Get("server.server_name").(string),
		"package user":             "package " + viper.Get("server.server_name").(string),
		"entityUser":               "entity" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
		"dtoUser":                  "dto" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context2)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
