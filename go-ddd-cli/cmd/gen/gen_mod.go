package gen

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/spf13/viper"
)

func GenMod(req GenReq) {
	str := `module github.com/jettjia/GasPc

go 1.20

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gogf/gf/v2 v2.1.1
	github.com/golang/mock v1.6.0
	github.com/google/wire v0.5.0
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.3 // indirect
	github.com/nsqio/go-nsq v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/goconvey v1.7.2
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
)
`
	path := req.BaseDir + "/go.mod"
	context := gstr.ReplaceByMap(str, g.MapStrStr{
		"go-ddd-demo-multi-system": viper.Get("server.go_module").(string),
	})

	if err := gfile.PutContents(path, context); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
