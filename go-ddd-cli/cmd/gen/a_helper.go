package gen

import (
	"strings"

	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

// generateStructFieldForModel 获取字段解析后的 type所索要的类型
// 要处理的目标
// CateName       string `gorm:"type:varchar(32); not null; default:0;comment:分类名称;" json:"cate_name" ` // 分类名称
// ormTag, gorm:"type:varchar(32); not null; comment:分类名称;"
// jsonTag, json:"cate_name"`
func generateStructFieldForModel(field TableColumn) (colStr string) {
	var fieldName, typeName, ormTag, jsonTag, node string

	typeName = generateStructFieldTypeName(field)

	// 字段名称 如CategoryName
	fieldName = GetJsonTagFromCase(field.Field, "Camel")

	// jsonTag 如 json:"cate_name"
	jsonTag = `json:"` + field.Field + `"`

	// note 如 // 分类名称
	node = " // " + field.Comment

	// ormTag 如 gorm:"column:category_name; type:varchar(32); not null; default:0; comment:分类名称;"
	// 如果是 created_at, updated_at, deleted_at 需要加毫秒
	if field.Field == "created_at" {
		ormTag = `gorm:"column:` + field.Field + ";autoCreateTime:milli;"
	} else if field.Field == "updated_at" {
		ormTag = `gorm:"column:` + field.Field + ";autoUpdateTime:milli;"
	} else if field.Field == "deleted_at" {
		ormTag = `gorm:"column:` + field.Field + ";autoDeletedTime:milli;"
	} else {
		ormTag = `gorm:"column:` + field.Field + ";"
	}

	//ormTag = `gorm:"`

	if gstr.ContainsI(field.Key, "pri") {
		ormTag += "primary_key;"
	}

	if gstr.ContainsI(field.Key, "pri") {
		if field.Field == "ulid" {
			ormTag += "type:" + field.Type
		} else {
			ormTag += "type:" + field.Type + " auto_increment;"
		}
	} else {
		ormTag += "type:" + field.Type + ";"
	}

	if gstr.ContainsI(field.Key, "uni") {
		ormTag += " uniqueIndex;"
	}

	if gstr.ContainsI(field.Key, "index") {
		ormTag += " index;"
	}

	if field.Null == "NO" {
		ormTag += "not null;"
	}

	if field.Default != "" {
		ormTag += "default: " + field.Default + ";"
	}

	if field.Comment != "" {
		ormTag += "comment:" + field.Comment + ";"
	}

	ormTag += `"`

	colStr = fieldName + "    " + typeName + "    " + "`" + ormTag + " " + jsonTag + "`" + node

	return
}

func generateStructFieldTypeName(field TableColumn) string {
	var typeName string
	t, _ := gregex.ReplaceString(`\(.+\)`, "", field.Type)
	t = gstr.Split(gstr.Trim(t), " ")[0]
	t = gstr.ToLower(t)
	switch t {
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		typeName = "[]byte"

	case "bit", "int", "int2", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "serial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint"
		} else {
			typeName = "int"
		}

	case "int4", "int8", "big_int", "bigint", "bigserial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint64"
		} else if field.Field == "deleted_at" {
			typeName = "int64"
		} else {
			typeName = "int64"
		}
	case "real":
		typeName = "float32"

	case "float", "double", "decimal", "smallmoney", "numeric":
		typeName = "float64"

	case "bool":
		typeName = "bool"

	case "datetime", "timestamp", "date", "time":
		typeName = "*time.Time"
	case "json":
		typeName = "string"
	default:
		// Automatically detect its data type.
		switch {
		case strings.Contains(t, "int"):
			typeName = "int"
		case strings.Contains(t, "text") || strings.Contains(t, "char"):
			typeName = "string"
		case strings.Contains(t, "float") || strings.Contains(t, "double"):
			typeName = "float64"
		case strings.Contains(t, "bool"):
			typeName = "bool"
		case strings.Contains(t, "binary") || strings.Contains(t, "blob"):
			typeName = "[]byte"
		case strings.Contains(t, "date") || strings.Contains(t, "time"):
			typeName = "*time.Time"
		default:
			typeName = "string"
		}
	}

	return typeName
}

func GetJsonTagFromCase(str, caseStr string) string {
	switch gstr.ToLower(caseStr) {
	case gstr.ToLower("Camel"): // 大驼峰
		return gstr.CaseCamel(str)

	case gstr.ToLower("CamelLower"): // 小驼峰
		return gstr.CaseCamelLower(str)

	case gstr.ToLower("Kebab"):
		return gstr.CaseKebab(str)

	case gstr.ToLower("KebabScreaming"):
		return gstr.CaseKebabScreaming(str)

	case gstr.ToLower("Snake"): // 将字符串转换中的符号(下划线,空格,点,中横线)用下划线( _ )替换,并全部转换为小写字母。
		return gstr.CaseSnake(str)

	case gstr.ToLower("SnakeFirstUpper"):
		return gstr.CaseSnakeFirstUpper(str)

	case gstr.ToLower("SnakeScreaming"):
		return gstr.CaseSnakeScreaming(str)
	case gstr.ToLower("ToLower"): // 全部小写
		return gstr.ToLower(str)
	}
	return str
}

// generateStructFieldForModel 获取字段解析后的 type所索要的类型
func generateStructFieldForEntity(field TableColumn) (colStr string) {
	var fieldName, typeName, jsonTag, node string

	typeName = generateStructFieldTypeNameForEntity(field)

	// 字段名称 如CategoryName
	fieldName = GetJsonTagFromCase(field.Field, "Camel")

	// jsonTag 如 json:"cate_name"
	jsonTag = `json:"` + field.Field + `"`

	// note 如 // 分类名称
	if field.Comment != "" {
		node = " // " + field.Comment
	}

	colStr = fieldName + "    " + typeName + "    " + "`" + jsonTag + "`" + node

	return
}

func generateStructFieldTypeNameForEntity(field TableColumn) string {
	var typeName string
	t, _ := gregex.ReplaceString(`\(.+\)`, "", field.Type)
	t = gstr.Split(gstr.Trim(t), " ")[0]
	t = gstr.ToLower(t)
	switch t {
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		typeName = "[]byte"

	case "bit", "int", "int2", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "serial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint"
		} else {
			typeName = "int"
		}

	case "int4", "int8", "big_int", "bigint", "bigserial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint64"
		} else {
			typeName = "int64"
		}

	case "real":
		typeName = "float32"

	case "float", "double", "decimal", "smallmoney", "numeric":
		typeName = "float64"

	case "bool":
		typeName = "bool"

	case "datetime", "timestamp", "date", "time":
		//typeName = "time.Time"
		typeName = "*gtime.Time"
	case "json":
		typeName = "string"
	default:
		// Automatically detect its data type.
		switch {
		case strings.Contains(t, "int"):
			typeName = "int"
		case strings.Contains(t, "text") || strings.Contains(t, "char"):
			typeName = "string"
		case strings.Contains(t, "float") || strings.Contains(t, "double"):
			typeName = "float64"
		case strings.Contains(t, "bool"):
			typeName = "bool"
		case strings.Contains(t, "binary") || strings.Contains(t, "blob"):
			typeName = "[]byte"
		case strings.Contains(t, "date") || strings.Contains(t, "time"):
			//typeName = "time.Time"
			typeName = "*gtime.Time"
		default:
			typeName = "string"
		}
	}

	return typeName
}

func genFieldForProtoMessage(field TableColumn) (colStr string) {
	var typeName string

	t, _ := gregex.ReplaceString(`\(.+\)`, "", field.Type)
	t = gstr.Split(gstr.Trim(t), " ")[0]
	t = gstr.ToLower(t)

	switch t {
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		typeName = "bytes"
	case "bit", "int", "int2", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "serial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint32"
		} else {
			typeName = "int32"
		}
	case "int4", "int8", "big_int", "bigint", "bigserial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint64"
		} else {
			typeName = "int64"
		}
	case "real":
		typeName = "float"
	case "float", "double", "decimal", "smallmoney", "numeric":
		typeName = "float"
	case "bool":
		typeName = "bool"
	case "datetime", "timestamp", "date", "time":
		typeName = "google.protobuf.Timestamp"
	case "varchar", "longtext", "text":
		typeName = "string"
	case "json":
		typeName = "string"
	}

	return typeName
}
