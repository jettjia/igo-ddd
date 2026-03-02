package gen

const TempGrpcCommonPb = `
syntax = "proto3";
option go_package = ".;proto";

message Empty {}

message Query {
    string key = 1; //表字段名称
    string value = 2; //表字段值
    Operator operator = 3; //判断条件
}

enum Operator {
    Operator_opEq          = 0;  // =
    Operator_opNe1         = 1;  // !=
    Operator_opNe2         = 2;  // <>
    Operator_opIn          = 3;  // in
    Operator_opNotIn       = 4;  // not in
    Operator_opGt          = 5;  // >
    Operator_opGte         = 6;  // >=
    Operator_opLt          = 7;  // <
    Operator_opLte         = 8;  // <=
    Operator_opLike        = 9;  // like
    Operator_opNotLike     = 10; // not like
    Operator_opBetween     = 11; // between
    Operator_opNotBetween  = 12; // not between
    Operator_opNull        = 13; // null
}

message PageData {
    uint32 page_num = 1; // 页码
    uint32 page_size = 2; // 每页显示行数
    uint32 total_number = 3; // 共多少条
    uint32 total_page = 4; // 共多少页
}

message SortData {
    string sort = 1; // 排序字段
    string direction = 2; // asc：升序;desc：降序
}
`

const TempGrpcServerPb = `
syntax = "proto3";
import "common.proto";
import "sys_menu.proto";

option go_package = ".;proto";

service Goods {
    // sys_menu
    rpc CreateSysMenu (CreateSysMenuReq) returns (CreateSysMenuRsp); // 创建
    rpc DeleteSysMenu (DeleteSysMenuReq) returns (Empty); // 删除
    rpc UpdateSysMenu (UpdateSysMenuReq) returns (Empty); // 修改
    rpc FindSysMenuById (FindSysMenuByIdReq) returns (FindSysMenuRsp); // 根据id查找
    rpc FindSysMenuPage (FindSysMenuPageReq) returns (FindSysMenuPageRsp); // 分页
}
`

const TempGrpcBizPb = `
syntax = "proto3";
option go_package = ".;proto";
import "common.proto";

message CreateSysMenuReq {
    {{CreateSysMenuReq}}
}

message CreateSysMenuRsp {
    string ulid = 1;
}

message DeleteSysMenuReq {
     string ulid = 1;
}

message UpdateSysMenuReq {
    {{UpdateSysMenuReq}}
}

message FindSysMenuByIdReq {
     string ulid = 1;
}

message FindSysMenuRsp {
	{{FindSysMenuRsp}}
}

message FindSysMenuPageReq{
    repeated Query query = 1;
    PageData page_data = 2;
    SortData sort_data = 3;
}

message FindSysMenuPageRsp {
    repeated  FindSysMenuRsp entries = 1;
    PageData page_data = 2;
}

`
