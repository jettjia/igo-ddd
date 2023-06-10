package sys

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"

	dto "github.com/jettjia/go-ddd-demo/application/dto/sys"
	service "github.com/jettjia/go-ddd-demo/application/service/sys"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/validate"
)

type sysMenuHandler struct {
	SysMenuSrv *service.SysMenuService
}

func NewSysMenuHandler() *sysMenuHandler {
	return &sysMenuHandler{
		SysMenuSrv: service.NewSysMenuService(),
	}
}

func (h *sysMenuHandler) CreateSysMenu(c *gin.Context) {
	// 参数解析
	dtoReq := dto.CreateSysMenuReq{}
	err := c.BindJSON(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	res, err := h.SysMenuSrv.CreateSysMenu(c.Request.Context(), &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusCreated, res.Id)
}

func (h *sysMenuHandler) DeleteSysMenu(c *gin.Context) {
	// 参数解析
	dtoReq := dto.DelSysMenusReq{}
	err := c.ShouldBindUri(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	err = h.SysMenuSrv.DeleteSysMenu(c.Request.Context(), &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusNoContent, nil)
}

func (h *sysMenuHandler) UpdateSysMenu(c *gin.Context) {
	// 参数解析
	dtoReq := dto.UpdateSysMenuReq{}
	err := c.ShouldBindUri(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}
	err = c.ShouldBindJSON(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}
	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	err = h.SysMenuSrv.UpdateSysMenu(c.Request.Context(), &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusNoContent, nil)
}

func (h *sysMenuHandler) FindSysMenuById(c *gin.Context) {
	// 参数解析
	dtoReq := dto.FindSysMenuByIdReq{}
	err := c.ShouldBindUri(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuById(c.Request.Context(), &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusOK, rsp)
}

func (h *sysMenuHandler) FindSysMenuByQuery(c *gin.Context) {
	// 参数解析
	dtoReq := dto.FindSysMenuByQueryReq{}
	err := c.ShouldBindJSON(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuByQuery(c.Request.Context(), &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusOK, rsp)
}

func (h *sysMenuHandler) FindSysMenuAll(c *gin.Context) {
	// 参数解析
	dtoReq := dto.FindSysMenuAllReq{}
	err := c.ShouldBindJSON(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuAll(c.Request.Context(), &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}
	responseutil.RspOk(c, http.StatusOK, rsp)
}

func (h *sysMenuHandler) FindSysMenuPage(c *gin.Context) {
	// 参数解析
	dtoReq := dto.FindSysMenuPageReq{}
	err := c.ShouldBindJSON(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuPage(c.Request.Context(), &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusOK, rsp)
}
