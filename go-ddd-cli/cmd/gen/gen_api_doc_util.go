package gen

import (
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

func utlGenPropertiesSon(col TableColumn) (str string) {
	str += `"` + col.Field + `": {`

	str += `"description": "` + col.Comment + `",`
	str += ` "type": "` + utilGenStructFieldTypeNameStr(col) + `"`

	str += `},`

	return
}

func utilGenRequiredSon(req GenReq) (colStr []string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "uuid" || v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" || v.Field == "created_by" || v.Field == "updated_by" || v.Field == "deleted_by" {
			continue
		}
		if gstr.ContainsI(v.Key, "pri") {
			continue
		}
		colStr = append(colStr, v.Field)
	}

	return colStr
}

func utilGenStructFieldTypeNameStr(field TableColumn) string {
	var typeName string
	t, _ := gregex.ReplaceString(`\(.+\)`, "", field.Type)
	t = gstr.Split(gstr.Trim(t), " ")[0]
	t = gstr.ToLower(t)
	switch t {
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		typeName = "[]byte"

	case "bit", "int", "int2", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "serial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "integer"
		} else {
			typeName = "integer"
		}

	case "int4", "int8", "big_int", "bigint", "bigserial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "integer"
		} else {
			typeName = "integer"
		}

	case "real":
		typeName = "number"

	case "float", "double", "decimal", "smallmoney", "numeric":
		typeName = "number"

	case "bool":
		typeName = "boolean"

	case "datetime", "timestamp", "date", "time":
		typeName = "string"
	case "json":
		typeName = "string"
	default:
		// Automatically detect its data type.
		switch {
		case strings.Contains(t, "int"):
			typeName = "integer"
		case strings.Contains(t, "text") || strings.Contains(t, "char"):
			typeName = "string"
		case strings.Contains(t, "float") || strings.Contains(t, "double"):
			typeName = "number"
		case strings.Contains(t, "bool"):
			typeName = "boolean"
		case strings.Contains(t, "binary") || strings.Contains(t, "blob"):
			typeName = "string"
		case strings.Contains(t, "date") || strings.Contains(t, "time"):
			typeName = "string"
		default:
			typeName = "string"
		}
	}

	return typeName
}

func utilGenExample(field TableColumn) (str string) {
	value := utilExampleValue(utilGenStructFieldTypeNameStr(field))
	str += `"` + field.Field + `": ` + gconv.String(value) + `,`

	return str
}

func utilExampleValue(t string) (value interface{}) {
	switch t {
	case "integer":
		value = 1
	case "number":
		value = 1.0
	case "boolean":
		value = true
	default:
		value = `"string"`
	}

	return value
}

func utilGenRequiredSonForQuery(req GenReq) (colStr []string) {
	if len(req.TableColumns) == 0 {
		return
	}

	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" || v.Field == "deleted_by" {
			continue
		}
		colStr = append(colStr, v.Field)
	}

	return colStr
}
