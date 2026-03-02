package gen

import (
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/spf13/viper"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// GenModel 生成po
func GenModel(req GenReq) {
	columnStr := "\n"

	for _, v := range req.TableColumns {
		colStr := generateStructFieldForModel(v)

		columnStr += colStr + "\n"
	}

	str := `package user` + `
import (
	"gorm.io/gorm"
	"e.coding.net/jidi/xdata/common/pkg/util"
)
type ` + GetJsonTagFromCase(req.TableName, "Camel") + ` struct {
	` + columnStr +
		`}

func (po *SysMenu) BeforeCreate(tx *gorm.DB) (err error) {
	po.Ulid = util.Ulid()
	return
}

func (po *SysMenu) TableName() string {
	return "sys_menu"
}
`

	str = gstr.ReplaceByMap(str, g.MapStrStr{
		"SysMenu":      GetJsonTagFromCase(req.TableName, "Camel"),
		"Ulid":         GetJsonTagFromCase(req.TableKey, "Camel"),  // 主键
		"sys_menu":     GetJsonTagFromCase(req.TableName, "Snake"), // 表名
		"package user": "package " + viper.Get("server.server_name").(string),
	})

	path := req.PoDir + "/" + req.TableName + "_po.go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

// 获取主键id的类型
func GetPriIdType(req GenReq) string {
	for _, v := range req.TableColumns {
		if gstr.ContainsI(v.Key, "pri") {
			return generateStructFieldTypeName(v)
		}
	}

	return ""
}
