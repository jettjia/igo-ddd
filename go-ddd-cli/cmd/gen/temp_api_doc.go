package gen

const TempApiDoc = `
{
    "openapi": "3.0.2",
    "info": {
        "version": "1.0.0",
        "title": "api3文档规范",
        "description": "API to access JettJia\r\n\r\n如有任何疑问，可到开发者社区提问：https://www.github.com/jettjia\r\n# Authentication\r\n- 调用需要鉴权的API，必须将token放在HTTP header中：\"Authorization: Bearer ACCESS_TOKEN\"\r\n",
        "contact": {},
        "x-logo": {
            "url": "data:image/png;base64,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",
            "backgroundColor": "#FFFFFF",
            "altText": "JettJia logo"
        }
    },
    "servers": [
        {
            "url": "https://{host}:{port}",
            "description": "host: JettJia服务器IP, port: 默认端口443",
            "variables": {
                "host": {
                    "default": "host",
                    "description": "JettJia服务器IP"
                },
                "port": {
                    "default": "443",
                    "description": "https默认端口443"
                }
            }
        }
    ],
    "paths": {
        "/api/pc/v1/sys/user": {
            "post": {
                "tags": [
                    "系统用户"
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "$ref": "#/components/schemas/CreateUser"
                            }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "201": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Id"
                                },
                                "examples": {
                                    "example": {
                                        "value": {
                                            "id": "01HD963WF3AXECGWSZDJR07W11"
                                        }
                                    }
                                }
                            }
                        },
                        "description": "创建成功 (返回创建成功的ulid)"
                    },
                    "400": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Error"
                                },
                                "examples": {
                                    "Err400000000": {
                                        "$ref": "#/components/examples/Err400000000"
                                    }
                                }
                            }
                        },
                        "description": "非法请求"
                    },
                    "401": {
                        "$ref": "#/components/responses/RespStd401"
                    },
                    "403": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Error"
                                },
                                "examples": {
                                    "Err403000000": {
                                        "$ref": "#/components/examples/Err403000000"
                                    }
                                }
                            }
                        },
                        "description": "禁止"
                    },
                    "409": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Error"
                                },
                                "examples": {
                                    "Err409000000": {
                                        "$ref": "#/components/examples/Err409000000"
                                    }
                                }
                            }
                        },
                        "description": "冲突"
                    },
                    "500": {
                        "$ref": "#/components/responses/RespStd500"
                    }
                },
                "summary": "创建用户"
            }
        },
        "/api/pc/v1/sys/user/{ulid}": {
            "get": {
                "tags": [
                    "系统用户"
                ],
                "parameters": [
                    {
                        "name": "ulid",
                        "description": "ulid",
                        "schema": {
                            "type": "string"
                        },
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/UserInfo"
                                }
                            }
                        },
                        "description": "接口调用成功"
                    },
                    "400": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Error"
                                },
                                "examples": {
                                    "Err400000000": {
                                        "$ref": "#/components/examples/Err400000000"
                                    }
                                }
                            }
                        },
                        "description": "非法请求"
                    },
                    "401": {
                        "$ref": "#/components/responses/RespStd401"
                    },
                    "403": {
                        "$ref": "#/components/responses/RespStd403"
                    },
                    "404": {
                        "$ref": "#/components/responses/RespStd404"
                    },
                    "500": {
                        "$ref": "#/components/responses/RespStd500"
                    }
                },
                "summary": "获取指定的用户"
            },
            "delete": {
                "tags": [
                    "系统用户"
                ],
                "parameters": [
                    {
                        "name": "ulid",
                        "description": "ulid",
                        "schema": {
                            "type": "string"
                        },
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "接口调用成功"
                    },
                    "400": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Error"
                                },
                                "examples": {
                                    "Err400000000": {
                                        "$ref": "#/components/examples/Err400000000"
                                    }
                                }
                            }
                        },
                        "description": "非法请求"
                    },
                    "401": {
                        "$ref": "#/components/responses/RespStd401"
                    },
                    "403": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Error"
                                },
                                "examples": {
                                    "Err403000000": {
                                        "$ref": "#/components/examples/Err403000000"
                                    }
                                }
                            }
                        },
                        "description": "禁止"
                    },
                    "404": {
                        "$ref": "#/components/responses/RespStd404"
                    },
                    "500": {
                        "$ref": "#/components/responses/RespStd500"
                    }
                },
                "summary": "删除用户"
            },
            "put": {
                "tags": [
                    "系统用户"
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "$ref": "#/components/schemas/UpdateUser"
                            }
                        }
                    },
                    "required": true
                },
                "parameters": [
                    {
                        "name": "ulid",
                        "description": "ulid",
                        "schema": {
                            "type": "string"
                        },
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "接口调用成功"
                    },
                    "400": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Error"
                                },
                                "examples": {
                                    "Err400000000": {
                                        "$ref": "#/components/examples/Err400000000"
                                    }
                                }
                            }
                        },
                        "description": "非法请求"
                    },
                    "401": {
                        "$ref": "#/components/responses/RespStd401"
                    },
                    "403": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Error"
                                },
                                "examples": {
                                    "Err403000000": {
                                        "$ref": "#/components/examples/Err403000000"
                                    }
                                }
                            }
                        },
                        "description": "禁止"
                    },
                    "404": {
                        "$ref": "#/components/responses/RespStd404"
                    },
                    "500": {
                        "$ref": "#/components/responses/RespStd500"
                    }
                },
                "summary": "编辑用户"
            }
        },
        "/api/pc/v1/sys/userPage": {
            "post": {
                "tags": [
                    "系统用户"
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "$ref": "#/components/schemas/Search"
                            }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/UserInfoPage"
                                }
                            }
                        },
                        "description": "接口调用成功"
                    },
                    "400": {
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Error"
                                },
                                "examples": {
                                    "Err400000000": {
                                        "$ref": "#/components/examples/Err400000000"
                                    }
                                }
                            }
                        },
                        "description": "非法请求"
                    },
                    "401": {
                        "$ref": "#/components/responses/RespStd401"
                    },
                    "403": {
                        "$ref": "#/components/responses/RespStd403"
                    },
                    "500": {
                        "$ref": "#/components/responses/RespStd500"
                    }
                },
                "summary": "获取用户列表",
                "description": ""
            }
        }
    },
    "components": {
        "schemas": {
            "Error": {
                "description": "如果接口调用返回的http状态码为非200系列，则表示发生异常，会返回错误码信息，具体参见错误码说明章节",
                "required": [
                    "code",
                    "message",
                    "cause"
                ],
                "type": "object",
                "properties": {
                    "code": {
                        "format": "int64",
                        "description": "业务错误码，前三位为 HTTP标准状态码，后三位为自定义状态码",
                        "type": "integer"
                    },
                    "message": {
                        "description": "可以直接展示给终端用户的错误信息\r\n",
                        "type": "string"
                    },
                    "cause": {
                        "description": "供开发者查看的错误信息\r\n",
                        "type": "string"
                    },
                    "detail": {
                        "description": "错误详细信息",
                        "type": "object",
                        "additionalProperties": true
                    }
                }
            },
			"SortData": {
                "description": "排序信息",
                "required": [
                    "sort",
                    "direction"
                ],
                "properties": {
                    "sort": {
                        "description": "排序字段，默认是id",
                        "type": "integer"
                    },
                    "direction": {
                        "description": "排序方式",
                        "type": "string",
                        "enum": [
                            "desc",
                            "asc"
                        ]
                    }
                },
                "example": {
                    "page_num": 1,
                    "page_size": 10,
                    "total_number": 9,
                    "total_page": 1
                }
            },
            "PageData": {
                "description": "分页信息",
                "required": [
                    "page_num",
                    "page_size"
                ],
                "properties": {
                    "page_num": {
                        "description": "当前页",
                        "type": "integer"
                    },
                    "page_size": {
                        "description": "每页数",
                        "type": "integer"
                    },
                    "total_number": {
                        "description": "总条数",
                        "type": "integer"
                    },
                    "total_page": {
                        "description": "总页数",
                        "type": "integer"
                    }
                },
                "example": {
                    "page_num": 1,
                    "page_size": 10,
                    "total_number": 9,
                    "total_page": 1
                }
            },
            "SearchQuery": {
                "description": "搜索字段说明",
                "required": [
                    "key",
                    "value",
                    "operator"
                ],
                "properties": {
                    "key": {
                        "description": "搜索的字段",
                        "type": "string"
                    },
                    "value": {
                        "description": "搜索的值",
                        "type": "string"
                    },
                    "operator": {
                        "description": "搜索的类型,比如等于、不等于、模糊搜索等;具体的值关系参考enum",
                        "type": "integer",
						"enum": [
                            "0:=",
                            "1:!=",
                            "2:<>",
                            "3:in",
                            "4:not in",
                            "5:>",
                            "6:>=",
                            "7:<",
                            "8:<=",
                            "9:like",
                            "10:not like",
                            "11:between",
                            "12:not between",
                            "13:null"
                        ]
                    }
                },
                "example": {
                    "key": "menu_name",
                    "value": "menu_name112",
                    "operator": 1
                }
            },
            "Search": {
                "description": "搜索",
                "required": [
                    "query",
                    "page_data",
                    "sort_data"
                ],
                "type": "object",
                "properties": {
                    "query": {
                        "$ref": "#/components/schemas/SearchQuery"
                    },
                    "page_data": {
                        "description": "分页;page_num:当前页;page_size:每页条数",
                        "type": "object",
						"$ref": "#/components/schemas/PageData"
                    },
                    "sort_data": {
                        "description": "排序;sort:排序字段;direction:排序类型,比如desc、asc;",
                        "type": "object",
						"$ref": "#/components/schemas/SortData"
                    }
                },
                "example": {
                    "query": [
                        {
                            "key": "menu_name",
                            "value": "menu_name112",
                            "operator": 1
                        }
                    ],
                    "page_data": {
                        "page_num": 1,
                        "page_size": 10
                    },
                    "sort_data": {
                        "sort": "ulid",
                        "direction": "desc"
                    }
                }
            },
            "Id": {
                "description": "id",
                "required": [
                    "ulid"
                ],
                "type": "object",
                "properties": {
                    "ulid": {
                        "description": "唯一标记Id",
                        "type": "string"
                    }
                },
                "example": {
                    "ulid": "01HD963WF3AXECGWSZDJR07W11"
                }
            },
            "UserBase": {
                "description": "用户信息基础",
                "required": [
                    {{UserBaseRequiredTable}}
                ],
                "type": "object",
                "properties": {
                    {{UserBaseProperties}}
                },
                "example": {
                     {{UserBaseExample}}
                }
            },
			"UserBaseInfo": {
                "description": "用户信息基础,查询",
                "required": [
                    {{UserBaseInfoRequiredTable}}
                ],
                "type": "object",
                "properties": {
                    {{UserBaseInfoProperties}}
                },
                "example": {
                     {{UserBaseInfoExample}}
                }
            },
            "UserInfo": {
                "description": "用户信息",
                "type": "object",
                "allOf": [
                    {
                        "$ref": "#/components/schemas/UserBaseInfo"
                    }
                ],
                "example": {
                    {{UserInfoExample}}
                }
            },
            "UserInfoPage": {
                "description": "用户列表",
                "required": [
                    "entries",
                    "page_data"
                ],
                "type": "object",
                "properties": {
                    "entries": {
                        "description": "用户信息",
                        "type": "array",
                        "items": {
                            "$ref": "#/components/schemas/UserInfo"
                        }
                    },
                    "page_data": {
                        "description": "分页信息",
                        "$ref": "#/components/schemas/PageData",
                        "type": "object"
                    }
                },
                "example": {
                    "entries": [
                        {
                             {{UserInfoPageExample}}
                        }
                    ],
                    "page_data": {
                        "page_num": 1,
                        "page_size": 10,
                        "total_number": 9,
                        "total_page": 1
                    }
                }
            },
            "CreateUser": {
                "description": "创建用户信息",
                "required": [
                    {{CreateUserRequiredTable}}
                ],
                "type": "object",
                "properties": {
                   {{CreateUserProperties}}
                },
                "example": {
                    {{CreateUserExample}}
                }
            },
            "UpdateUser": {
                "description": "编辑用户信息",
                "required": [
                    "id"
                ],
                "type": "object",
                "properties": {
                    {{UpdateUserProperties}}
                },
                "example": {
                    {{UpdateUserExample}}
                }
            }
        },
        "responses": {
            "RespStd400": {
                "content": {
                    "application/json": {
                        "schema": {
                            "$ref": "#/components/schemas/Error"
                        },
                        "examples": {
                            "Err400000000": {
                                "$ref": "#/components/examples/Err400000000"
                            }
                        }
                    }
                },
                "description": "非法请求"
            },
            "RespStd401": {
                "content": {
                    "application/json": {
                        "schema": {
                            "$ref": "#/components/schemas/Error"
                        },
                        "examples": {
                            "Err401000000": {
                                "$ref": "#/components/examples/Err401000000"
                            }
                        }
                    }
                },
                "description": "未授权"
            },
            "RespStd403": {
                "content": {
                    "application/json": {
                        "schema": {
                            "$ref": "#/components/schemas/Error"
                        },
                        "examples": {
                            "Err403000000": {
                                "$ref": "#/components/examples/Err403000000"
                            }
                        }
                    }
                },
                "description": "无法执行此操作"
            },
            "RespStd404": {
                "content": {
                    "application/json": {
                        "schema": {
                            "$ref": "#/components/schemas/Error"
                        },
                        "examples": {
                            "Err404000000": {
                                "$ref": "#/components/examples/Err404000000"
                            }
                        }
                    }
                },
                "description": "资源不存在"
            },
            "RespStd417": {
                "content": {
                    "application/json": {
                        "schema": {
                            "$ref": "#/components/schemas/Error"
                        },
                        "examples": {
                            "Err403000000": {
                                "$ref": "#/components/examples/Err417000000"
                            }
                        }
                    }
                },
                "description": "检测的状态没通过"
            },
            "RespStd500": {
                "content": {
                    "application/json": {
                        "schema": {
                            "$ref": "#/components/schemas/Error"
                        },
                        "examples": {
                            "Err500000000": {
                                "$ref": "#/components/examples/Err500000000"
                            }
                        }
                    }
                },
                "description": "内部错误"
            }
        },
        "examples": {
            "Err400000000": {
                "summary": "参数错误",
                "description": "以下原因均会导致该错误\n- 未提供必须参数\n- 参数类型错误\n- 参数结构不正确\n- 枚举不存在\n- 参数值不符合要求\n",
                "value": {
                    "cause": "...",
                    "code": 400100000,
                    "message": "非法参数"
                }
            },
            "Err401000000": {
                "summary": "授权无效",
                "value": {
                    "cause": "...",
                    "code": 401100000,
                    "message": "授权无效"
                }
            },
            "Err403000000": {
                "summary": "没有权限执行此操作",
                "value": {
                    "cause": "...",
                    "code": 403100000,
                    "message": "没有权限执行此操作"
                }
            },
            "Err404000000": {
                "summary": "记录不存在",
                "value": {
                    "cause": "...",
                    "code": 404100000,
                    "message": "记录不存在"
                }
            },
            "Err409000000": {
                "summary": "已存在关联的相关信息",
                "value": {
                    "cause": "...",
                    "code": 409100000,
                    "message": "已存在关联的相关信息"
                }
            },
            "Err417000000": {
                "summary": "校验状态失败",
                "value": {
                    "cause": "...",
                    "code": 409100000,
                    "message": "校验状态失败"
                }
            },
            "Err500000000": {
                "summary": "内部错误",
                "value": {
                    "cause": "...",
                    "code": 500100000,
                    "message": "内部错误"
                }
            }
        },
        "securitySchemes": {
            "OAuth2.0": {
                "flows": {},
                "type": "oauth2"
            }
        }
    },
    "security": [
        {
            "OAuth2.0": []
        }
    ],
    "tags": [
        {
            "name": "系统用户",
            "description": ""
        }
    ]
}
`
