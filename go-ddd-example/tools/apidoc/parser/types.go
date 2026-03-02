package parser

// Config 解析器配置
type Config struct {
	RouterPath  string // 路由文件路径
	HandlerPath string // Handler 文件路径
	DtoPath     string // DTO 文件路径
}

// Route 路由信息
type Route struct {
	Path       string // 路径
	Method     string // HTTP 方法
	Handler    string // Handler 函数名
	HandlerPkg string // Handler 包名
}

// Handler Handler 信息
type Handler struct {
	Name      string   // 函数名
	Package   string   // 包名
	Request   *DtoInfo // 请求 DTO
	Response  *DtoInfo // 响应 DTO
	Summary   string   // 接口摘要
	Desc      string   // 接口详细描述
	URLParams []string // URL 参数
}

// DtoInfo DTO 信息
type DtoInfo struct {
	Name    string          // 结构体名称
	Package string          // 包名
	Fields  []*DtoFieldInfo // 字段信息
}

// DtoFieldInfo DTO 字段信息
type DtoFieldInfo struct {
	Name     string // 字段名
	Type     string // 字段类型
	Tag      string // 字段标签
	Required bool   // 是否必填
	Comment  string // 字段注释
}
