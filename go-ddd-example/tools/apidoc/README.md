# 自动化接口文档

## 注释模式
使用注释模式，会覆盖掉自动模式生成的文档，优先级更高
```go
// @summary 查询用户详情
// @desc 这是接口的详细描述，支持多行内容。
// @request FindSysUserByIdReq
// @response FindSysUserRsp
func (h *Handler) FindSysUserById(c *gin.Context) {
    // ...
}
```