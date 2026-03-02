package gen

import (
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/spf13/viper"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/os/gfile"
)

// GenEntity 生成entity
func GenDomainEntity(req GenReq) {
	str := `package user

	import (
		"xdata/xtext/zsvc-example/infra/repository/po/user"
	)

	type SysMenu struct {
		user.SysMenu
	}
`

	context := gstr.ReplaceByMap(str, g.MapStrStr{
		"package user":             "package " + viper.Get("server.server_name").(string),
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"SysMenu":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"po/user":                  "po/" + viper.Get("server.server_name").(string),
		"user.":                    viper.Get("server.server_name").(string) + ".",
	})

	path := req.EntityDir + "/" + req.TableName + "_entity.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
