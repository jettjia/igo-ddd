package gen

const TempInterfacesHandler = `
package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jettjia/igo-pkg/pkg/validate"
	"github.com/jettjia/igo-pkg/pkg/xerror"
	"github.com/jettjia/igo-pkg/pkg/xresponse"

	dtoUser "xdata/xtext/zsvc-example/application/dto/user"
	"xdata/xtext/zsvc-example/types/apierror"
)

// @response CreateSysMenuRsp
func (h *Handler) CreateSysMenu(c *gin.Context) {
	// 参数解析
	dtoReq := dtoUser.CreateSysMenuReq{}
	err := c.BindJSON(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}
	dtoReq.CreatedBy = c.GetString("user_id")

	// 业务处理
	res, err := h.SysMenuSrv.CreateSysMenu(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	xresponse.RspOk(c, http.StatusCreated, res)
}

func (h *Handler) DeleteSysMenu(c *gin.Context) {
	// 参数解析
	dtoReq := dtoUser.DelSysMenusReq{}
	err := c.ShouldBindUri(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}
	dtoReq.DeletedBy = c.GetString("user_id")

	// 业务处理
	err = h.SysMenuSrv.DeleteSysMenu(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	xresponse.RspOk(c, http.StatusNoContent, nil)
}

func (h *Handler) UpdateSysMenu(c *gin.Context) {
	// 参数解析
	dtoReq := dtoUser.UpdateSysMenuReq{}
	err := c.ShouldBindUri(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}
	err = c.BindJSON(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}
	dtoReq.UpdatedBy = c.GetString("user_id")

	// 业务处理
	err = h.SysMenuSrv.UpdateSysMenu(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	xresponse.RspOk(c, http.StatusNoContent, nil)
}

// @response FindSysMenuRsp
func (h *Handler) FindSysMenuById(c *gin.Context) {
	// 参数解析
	dtoReq := dtoUser.FindSysMenuByIdReq{}
	err := c.ShouldBindUri(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuById(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	xresponse.RspOk(c, http.StatusOK, rsp)
}

// @response FindSysMenuRsp
func (h *Handler) FindSysMenuByQuery(c *gin.Context) {
	// 参数解析
	dtoReq := dtoUser.FindSysMenuByQueryReq{}
	err := c.BindJSON(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuByQuery(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	xresponse.RspOk(c, http.StatusOK, rsp)
}

// @response []FindSysMenuRsp
func (h *Handler) FindSysMenuAll(c *gin.Context) {
	// 参数解析
	dtoReq := dtoUser.FindSysMenuAllReq{}
	err := c.BindJSON(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuAll(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}
	xresponse.RspOk(c, http.StatusOK, rsp)
}

// @response FindSysMenuPageRsp
func (h *Handler) FindSysMenuPage(c *gin.Context) {
	// 参数解析
	dtoReq := dtoUser.FindSysMenuPageReq{}
	err := c.BindJSON(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(err.Error()))
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuPage(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	xresponse.RspOk(c, http.StatusOK, rsp)
}
`
