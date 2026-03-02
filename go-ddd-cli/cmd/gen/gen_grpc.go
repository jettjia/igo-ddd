package gen

import (
	"fmt"
	"strings"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/spf13/viper"
)

func GenGrpcCode(req GenReq) {
	// create pb
	genGrpcPb(req)

	// create handler
	genGrpcHandler(req)

	// create ginit
	genGrpcGinit(req)
}

func genGrpcPb(req GenReq) {
	genGrpcPbCommon(req)
	genGrpcPbServer(req)
	genGrpcPbBiz(req)
}

func genGrpcPbCommon(req GenReq) {
	context := gstr.ReplaceByMap(TempGrpcCommonPb, g.MapStrStr{})

	path := req.GrpcDir + "/proto/" + GetJsonTagFromCase(viper.Get("grpc.grpc_server").(string), "CamelLower") + "/common.proto"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genGrpcPbServer(req GenReq) {
	context := gstr.ReplaceByMap(TempGrpcServerPb, g.MapStrStr{
		`"Goods"`: `"` + viper.Get("grpc.grpc_server").(string) + `"`,
		`SysMenu`: GetJsonTagFromCase(req.TableName, "Camel"),
	})

	serverName := GetJsonTagFromCase(viper.Get("grpc.grpc_server").(string), "CamelLower")
	path := req.GrpcDir + "/proto/" + serverName + "/" + serverName + ".proto"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genGrpcPbBiz(req GenReq) {
	context := gstr.ReplaceByMap(TempGrpcBizPb, g.MapStrStr{
		`SysMenu`:              GetJsonTagFromCase(req.TableName, "Camel"),
		`{{CreateSysMenuReq}}`: _genGrpcPbBizCreateReq(req),
		`{{UpdateSysMenuReq}}`: _genGrpcPbBizUpdateReq(req),
		`{{FindSysMenuRsp}}`:   _genGrpcPbBizFindRsp(req),
	})

	serverName := GetJsonTagFromCase(viper.Get("grpc.grpc_server").(string), "CamelLower")
	path := req.GrpcDir + "/proto/" + serverName + "/" + GetJsonTagFromCase(req.TableName, "Snake") + ".proto"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

/*
*

		格式:
		string menuName = 1;
	    string desc = 2;
*/
func _genGrpcPbBizCreateReq(req GenReq) (str string) {

	count := 0
	for _, v := range req.TableColumns {
		if v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" || v.Field == "created_by" || v.Field == "updated_by" || v.Field == "deleted_by" {
			continue
		}
		if gstr.ContainsI(v.Key, "pri") {
			continue
		}

		typeName := genFieldForProtoMessage(v)
		count += 1

		if count == 1 {
			str += fmt.Sprintf("%s %s = %d; // %s", typeName, GetJsonTagFromCase(v.Field, "CamelLower"), count, v.Comment) + "\n"
		} else {
			str += "    " + fmt.Sprintf("%s %s = %d; // %s", typeName, GetJsonTagFromCase(v.Field, "CamelLower"), count, v.Comment) + "\n"
		}
	}

	return
}

func _genGrpcPbBizUpdateReq(req GenReq) (str string) {

	count := 0
	for _, v := range req.TableColumns {
		if v.Field == "created_at" || v.Field == "updated_at" || v.Field == "deleted_at" || v.Field == "created_by" || v.Field == "updated_by" || v.Field == "deleted_by" {
			continue
		}

		typeName := genFieldForProtoMessage(v)
		count += 1

		if count == 1 {
			str += fmt.Sprintf("%s %s = %d; // %s", typeName, GetJsonTagFromCase(v.Field, "CamelLower"), count, v.Comment) + "\n"
		} else {
			str += "    " + fmt.Sprintf("%s %s = %d; // %s", typeName, GetJsonTagFromCase(v.Field, "CamelLower"), count, v.Comment) + "\n"
		}
	}

	return
}

func _genGrpcPbBizFindRsp(req GenReq) (str string) {

	count := 0
	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" || v.Field == "updated_by" {
			continue
		}

		typeName := genFieldForProtoMessage(v)
		count += 1

		if count == 1 {
			str += fmt.Sprintf("%s %s = %d; // %s", typeName, GetJsonTagFromCase(v.Field, "CamelLower"), count, v.Comment) + "\n"
		} else {
			str += "    " + fmt.Sprintf("%s %s = %d; // %s", typeName, GetJsonTagFromCase(v.Field, "CamelLower"), count, v.Comment) + "\n"
		}
	}

	return
}

func genGrpcHandler(req GenReq) {
	genGrpcHandlerBase(req)
	genGrpcHandlerBiz(req)
}

func genGrpcHandlerBase(req GenReq) {
	context := gstr.ReplaceByMap(TempGrpcHandlerBase, g.MapStrStr{
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		`SysMenu`:                  GetJsonTagFromCase(req.TableName, "Camel"),
	})

	path := req.GrpcDir + "/ghandler/" + "base.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genGrpcHandlerBiz(req GenReq) {
	context := gstr.ReplaceByMap(TempGrpcHandlerBiz, g.MapStrStr{
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"grpcGoodsProto":           "grpc" + viper.Get("grpc.grpc_server").(string) + "Proto",
		"proto/goods":              "proto/" + GetJsonTagFromCase(viper.Get("grpc.grpc_server").(string), "CamelLower"),
		`SysMenu`:                  GetJsonTagFromCase(req.TableName, "Camel"),
	})

	path := req.GrpcDir + "/ghandler/" + "g_" + GetJsonTagFromCase(req.TableName, "Snake") + "_handler.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genGrpcGinit(req GenReq) {
	context := gstr.ReplaceByMap(TempGrpcGinit, g.MapStrStr{
		"xdata/xtext/zsvc-example": viper.Get("server.go_module").(string),
		"grpcGoodsProto":           "grpc" + viper.Get("grpc.grpc_server").(string) + "Proto",
		"proto/goods":              "proto/" + GetJsonTagFromCase(viper.Get("grpc.grpc_server").(string), "CamelLower"),
		`SysMenu`:                  GetJsonTagFromCase(req.TableName, "Camel"),
	})

	path := req.GrpcDir + "/ginit/g_init_grpc_server.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
