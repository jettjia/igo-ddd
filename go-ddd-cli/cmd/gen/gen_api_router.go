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

func GenInterfacesRouter(req GenReq) {
	context := gstr.ReplaceByMap(TempInterfacesRouter, g.MapStrStr{
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"handler/public/user":      "handler/public/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"Group(\"/user":            "Group(\"/" + viper.Get("server.server_name").(string),
		"SysMenu":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"SetPublicRouter":          "SetPublicRouter" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),

		`"/menu"`:      `"/` + GetJsonTagFromCase(req.TableName, "CamelLower") + `"`,     //创建
		`"/menu/:`:     `"/` + GetJsonTagFromCase(req.TableName, "CamelLower") + `/:`,    //删除、修改、查看
		`"/menu/by`:    `"/` + GetJsonTagFromCase(req.TableName, "CamelLower") + `/by`,   //删除、修改、查看
		`"/menuPage"`:  `"/` + GetJsonTagFromCase(req.TableName, "CamelLower") + `Page"`, // 查询分页
		"package user": "package " + viper.Get("server.server_name").(string),
		"handUser":     "hand" + GetJsonTagFromCase(req.SrvName, "Camel"),
		"// menu":      "// " + GetJsonTagFromCase(viper.Get("mysql.table").(string), "ToLower"),
	})

	path := req.RouterDir + "/" + req.TableName + "_router.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
