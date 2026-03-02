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

// GenConverter gen converter
func GenConverter(req GenReq) {
	path := req.ConverterDir + "/" + req.TableName + "_conv.go"

	context := gstr.ReplaceByMap(TempInfrastructureConverter, g.MapStrStr{
		"package user":             "package " + viper.Get("server.server_name").(string),
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"entityUser":               "entity" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
		"poUser":                   "po" + GetJsonTagFromCase(viper.Get("server.server_name").(string), "Camel"),
		"SysMenu":                  GetJsonTagFromCase(req.TableName, "Camel"),
		"{{createUuid}}":           genE2PAdd(req),
		"entity/user":              "entity/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		"po/user":                  "po/" + GetJsonTagFromCase(req.SrvName, "ToLower"),
		//"{{createConv}}":  genE2PAdd(req),
		//"{{updateConv}}":  genE2PUpdate(req),
		//"{{findOneConv}}": genE2PFindOne(req),
	})
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genE2PAdd(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	modelName := GetJsonTagFromCase(req.TableName, "CamelLower")

	for _, v := range req.TableColumns {
		if v.Field == "uuid" {
			return modelName + ".Uuid = util.Ulid()"
		}

	}

	return
}

func genE2PAddBak(req GenReq) (str string) {
	return
	if len(req.TableColumns) == 0 {
		return
	}

	modelName := GetJsonTagFromCase(req.TableName, "CamelLower")

	for _, v := range req.TableColumns {
		if v.Field == "updated_by" || v.Field == "deleted_at" || v.Field == "deleted_by" || v.Field == req.TableKey {
			continue
		}

		if v.Field == "updated_at" || v.Field == "created_at" {
			str += modelName + "." + GetJsonTagFromCase(v.Field, "Camel") + "= gtime.Now()" + "\n"
		} else {
			str += modelName + "." + GetJsonTagFromCase(v.Field, "Camel") + "= en." + GetJsonTagFromCase(v.Field, "Camel") + "\n"
		}
	}

	return
}

func genE2PUpdate(req GenReq) (str string) {
	return
	if len(req.TableColumns) == 0 {
		return
	}

	modelName := GetJsonTagFromCase(req.TableName, "CamelLower")

	for _, v := range req.TableColumns {
		if v.Field == "created_at" || v.Field == "created_by" || v.Field == "deleted_at" || v.Field == "deleted_by" || v.Field == req.TableKey {
			continue
		}
		if v.Field == "updated_at" {
			str += modelName + "." + GetJsonTagFromCase(v.Field, "Camel") + "= gtime.Now()" + "\n"
		} else {
			str += modelName + "." + GetJsonTagFromCase(v.Field, "Camel") + "= en." + GetJsonTagFromCase(v.Field, "Camel") + "\n"
		}
	}

	return
}

func genE2PFindOne(req GenReq) (str string) {
	return
	if len(req.TableColumns) == 0 {
		return
	}

	modelName := GetJsonTagFromCase(req.TableName, "CamelLower")

	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" || v.Field == "deleted_by" {
			continue
		}
		str += modelName + "." + GetJsonTagFromCase(v.Field, "Camel") + "= po." + GetJsonTagFromCase(v.Field, "Camel") + "\n"
	}

	return
}
