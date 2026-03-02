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

// GenApiDoc 自动化文档
func GenApiDoc(req GenReq) {
	context := gstr.ReplaceByMap(TempApiDoc, g.MapStrStr{
		// paths 路由转换
		`"系统用户"`:          `"` + viper.Get("doc.api_tag").(string) + `"`,
		`/pc/v1/sys/user`: "/" + viper.Get("doc.api_module").(string) + "/" + viper.Get("doc.api_version").(string) + "/" + viper.Get("doc.api_module_son").(string) + "/" + viper.Get("doc.api_table_short_name").(string),
		`"创建用户"`:          `"创建` + viper.Get("doc.api_table").(string) + `"`,
		`"获取指定的用户"`:       `"获取指定的` + viper.Get("doc.api_table").(string) + `"`,
		`"删除用户"`:          `"删除` + viper.Get("doc.api_table").(string) + `"`,
		`"编辑用户"`:          `"编辑` + viper.Get("doc.api_table").(string) + `"`,
		`"获取用户列表"`:        `"获取` + viper.Get("doc.api_table").(string) + `列表"`,

		// components/scheme转换，下面都是
		// base，创建、修改依赖
		`"UserBase"`:                `"` + GetJsonTagFromCase(viper.Get("doc.api_table_short_name").(string), "Camel") + `Base"`,
		"用户信息基础":                    viper.Get("doc.api_table").(string) + "信息基础",
		"{{UserBaseRequiredTable}}": genUserBaseRequiredSonStr(req),
		"{{UserBaseProperties}}":    genUserBaseProperties(req),
		"{{UserBaseExample}}":       genUserBaseExample(req),

		// base，查询依赖依赖
		`"UserBaseInfo"`:                `"` + GetJsonTagFromCase(viper.Get("doc.api_table_short_name").(string), "Camel") + `BaseInfo"`,
		"用户信息基础,查询":                     viper.Get("doc.api_table").(string) + "信息基础,查询",
		"{{UserBaseInfoRequiredTable}}": genUserBaseInfoRequiredSonStr(req),
		"{{UserBaseInfoProperties}}":    genUserBaseInfoProperties(req),
		"{{UserBaseInfoExample}}":       genUserBaseInfoExample(req),

		// info
		`"UserInfo"`:                   `"` + GetJsonTagFromCase(viper.Get("doc.api_table_short_name").(string), "Camel") + `Info"`,
		`"用户信息"`:                       `"` + viper.Get("doc.api_table").(string) + `信息"`,
		`/components/schemas/UserBase`: "/components/schemas/" + GetJsonTagFromCase(viper.Get("doc.api_table_short_name").(string), "Camel") + "Base",
		"{{UserInfoExample}}":          genUseInfoExample(req),

		// infos
		`"UserInfoPage"`:               `"` + GetJsonTagFromCase(viper.Get("doc.api_table_short_name").(string), "Camel") + `InfoPage"`,
		`"用户列表"`:                       `"` + viper.Get("doc.api_table").(string) + `列表"`,
		`/components/schemas/UserInfo`: `/components/schemas/` + GetJsonTagFromCase(viper.Get("doc.api_table_short_name").(string), "Camel") + `Info`,
		"{{UserInfoPageExample}}":      genUseInfosExample(req),

		// create
		`"CreateUser"`:                `"Create` + GetJsonTagFromCase(viper.Get("doc.api_table_short_name").(string), "Camel") + `"`,
		`创建用户信息`:                      `创建` + viper.Get("doc.api_table").(string) + `信息`,
		"{{CreateUserRequiredTable}}": genCreateUserRequiredSonStr(req),
		"{{CreateUserProperties}}":    genCreateUserProperties(req),
		"{{CreateUserExample}}":       genCreateUserExample(req),
		"/CreateUser":                 `/Create` + GetJsonTagFromCase(viper.Get("doc.api_table_short_name").(string), "Camel"),

		// update
		`"UpdateUser"`:             `"Update` + GetJsonTagFromCase(viper.Get("doc.api_table_short_name").(string), "Camel") + `"`,
		`编辑用户信息`:                   `编辑` + viper.Get("doc.api_table").(string) + `信息`,
		"{{UpdateUserProperties}}": genUpdateUserProperties(req),
		"{{UpdateUserExample}}":    genUpdateUserExample(req),
		"/UpdateUser":              `/Update` + GetJsonTagFromCase(viper.Get("doc.api_table_short_name").(string), "Camel"),
	})

	path := req.ApiDocDir + "/" + viper.Get("mysql.table").(string) + "_api_doc.json"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

// genRequiredSonStr 获取required字段信息
func genUserBaseRequiredSonStr(req GenReq) (respColStr string) {
	colStrArr := utilGenRequiredSon(req)
	colStrArrLen := len(colStrArr)
	if colStrArrLen == 0 {
		return
	}

	for k, v := range colStrArr {
		if k < colStrArrLen-1 {
			respColStr += `"` + v + `",`
		} else {
			respColStr += `"` + v + `"`
		}

	}

	return
}

// genProperties 生成properties
func genUserBaseProperties(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" || v.Field == "deleted_by" {
			continue
		}
		str += utlGenPropertiesSon(v)
	}

	str = strings.TrimRight(str, ",")

	return
}

// genExample
func genUserBaseExample(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" || v.Field == "deleted_by" {
			continue
		}
		str += utilGenExample(v)
	}

	str = strings.TrimRight(str, ",")

	return
}

// genExample
func genUseInfoExample(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" || v.Field == "deleted_by" {
			continue
		}
		str += utilGenExample(v)
	}
	str = strings.TrimRight(str, ",")

	return
}

// genExample
func genUseInfosExample(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" || v.Field == "deleted_by" {
			continue
		}
		str += utilGenExample(v)
	}
	str = strings.TrimRight(str, ",")

	return
}

// genProperties 生成properties
func genCreateUserRequiredSonStr(req GenReq) (respColStr string) {
	colStrArr := utilGenRequiredSon(req)
	colStrArrLen := len(colStrArr)
	if colStrArrLen == 0 {
		return
	}

	for k, v := range colStrArr {
		if k < colStrArrLen-1 {
			respColStr += `"` + v + `",`
		} else {
			respColStr += `"` + v + `"`
		}

	}

	return
}

// genCreateUserProperties 获取required字段信息
func genCreateUserProperties(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "id" || v.Field == "uuid" || v.Field == "ulid" || v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" || v.Field == "created_by" || v.Field == "updated_by" || v.Field == "deleted_by" {
			continue
		}
		str += utlGenPropertiesSon(v)
	}

	str = strings.TrimRight(str, ",")

	return
}

// genCreateUserExample
func genCreateUserExample(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {

		if v.Field == "uuid" || v.Field == "ulid" || v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" || v.Field == "created_by" || v.Field == "updated_by" || v.Field == "deleted_by" {
			continue
		}
		if gstr.ContainsI(v.Key, "pri") {
			continue
		}
		str += utilGenExample(v)
	}

	str = strings.TrimRight(str, ",")

	return
}

// genCreateUserProperties 获取required字段信息
func genUpdateUserProperties(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "uuid" || v.Field == "ulid" || v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" || v.Field == "created_by" || v.Field == "updated_by" || v.Field == "deleted_by" {
			continue
		}
		if gstr.ContainsI(v.Key, "pri") {
			continue
		}
		str += utlGenPropertiesSon(v)
	}

	str = strings.TrimRight(str, ",")

	return
}

// genExample
func genUpdateUserExample(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "uuid" || v.Field == "ulid" || v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" || v.Field == "created_by" || v.Field == "updated_by" || v.Field == "deleted_by" {
			continue
		}
		if gstr.ContainsI(v.Key, "pri") {
			continue
		}
		str += utilGenExample(v)
	}

	str = strings.TrimRight(str, ",")

	return
}

// genRequiredSonStr 获取required字段信息
func genUserBaseInfoRequiredSonStr(req GenReq) (respColStr string) {
	colStrArr := utilGenRequiredSonForQuery(req)
	colStrArrLen := len(colStrArr)
	if colStrArrLen == 0 {
		return
	}

	for k, v := range colStrArr {
		if k < colStrArrLen-1 {
			respColStr += `"` + v + `",`
		} else {
			respColStr += `"` + v + `"`
		}

	}

	return
}

// genProperties 生成properties
func genUserBaseInfoProperties(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" || v.Field == "deleted_by" {
			continue
		}
		str += utlGenPropertiesSon(v)
	}

	str = strings.TrimRight(str, ",")

	return
}

// genExample
func genUserBaseInfoExample(req GenReq) (str string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" || v.Field == "deleted_by" {
			continue
		}
		str += utilGenExample(v)
	}

	str = strings.TrimRight(str, ",")

	return
}
