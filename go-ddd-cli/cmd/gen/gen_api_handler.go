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

func GenInterfacesHandler(req GenReq) {
	context := gstr.ReplaceByMap(TempInterfacesHandler, g.MapStrStr{
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"dto/user":                 "dto/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"service/user":             "service/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"SysMenu":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"sys_menu":                 GetJsonTagFromCase(req.TableName, "Snake"), // 表名
		"MenuId":                   GetJsonTagFromCase(req.TableKey, "Camel"),  // 主键
		"package user":             "package " + viper.Get("server.server_name").(string),
		"dtoUser":                  "dto" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
	})

	path := req.HandlerDir + "/" + req.TableName + "_handler.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}

	genHandler(req)
}

func genHandler(req GenReq) {
	context := gstr.ReplaceByMap(TempHandler, g.MapStrStr{
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"service/user":             "service/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"serviceUser":              "service" + GetJsonTagFromCase(req.SrvName, "Camel"),
		"SysMenu":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"package user":             "package " + viper.Get("server.server_name").(string),
	})

	path := req.HandlerDir + "/" + "handler.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
