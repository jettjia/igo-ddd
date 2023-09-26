package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"

	"jettjia/go-ddd-demo-multi-common/pkg/response"
	"jettjia/go-ddd-demo-multi-common/pkg/validate"

	"jettjia/go-ddd-demo-multi-system/application/dto"
	"jettjia/go-ddd-demo-multi-system/domain/entity"
)

func (h *Handler) CreateSysMenu(c *gin.Context) {
	// 参数解析
	dtoReq := dto.CreateSysMenuReq{}
	err := c.BindJSON(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}
	tokenData, _ := entity.GinParse(c)
	if tokenData != nil {
		dtoReq.CreatedBy = tokenData.Username
	}

	// 业务处理
	res, err := h.SysMenuSrv.CreateSysMenu(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.RspOk(c, http.StatusCreated, res)
}

func (h *Handler) DeleteSysMenu(c *gin.Context) {
	// 参数解析
	dtoReq := dto.DelSysMenusReq{}
	err := c.ShouldBindUri(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}
	tokenData, _ := entity.GinParse(c)
	if tokenData != nil {
		dtoReq.DeletedBy = tokenData.Username
	}

	// 业务处理
	err = h.SysMenuSrv.DeleteSysMenu(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.RspOk(c, http.StatusNoContent, nil)
}

func (h *Handler) UpdateSysMenu(c *gin.Context) {
	// 参数解析
	dtoReq := dto.UpdateSysMenuReq{}
	err := c.ShouldBindUri(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}
	err = c.BindJSON(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}
	tokenData, _ := entity.GinParse(c)
	if tokenData != nil {
		dtoReq.UpdatedBy = tokenData.Username
	}

	// 业务处理
	err = h.SysMenuSrv.UpdateSysMenu(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.RspOk(c, http.StatusNoContent, nil)
}

func (h *Handler) FindSysMenuById(c *gin.Context) {
	// 参数解析
	dtoReq := dto.FindSysMenuByIdReq{}
	err := c.ShouldBindUri(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuById(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if rsp.Ulid == "" {
		err = gerror.NewCode(response.CommNotFound)
		_ = c.Error(err)
		return
	}

	response.RspOk(c, http.StatusOK, rsp)
}

func (h *Handler) FindSysMenuByQuery(c *gin.Context) {
	// 参数解析
	dtoReq := dto.FindSysMenuByQueryReq{}
	err := c.BindJSON(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuByQuery(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if rsp.Ulid == "" {
		err = gerror.NewCode(response.CommNotFound)
		_ = c.Error(err)
		return
	}

	response.RspOk(c, http.StatusOK, rsp)
}

func (h *Handler) FindSysMenuAll(c *gin.Context) {
	// 参数解析
	dtoReq := dto.FindSysMenuAllReq{}
	err := c.BindJSON(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuAll(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.RspOk(c, http.StatusOK, rsp)
}

func (h *Handler) FindSysMenuPage(c *gin.Context) {
	// 参数解析
	dtoReq := dto.FindSysMenuPageReq{}
	err := c.BindJSON(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(response.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuPage(c, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.RspOk(c, http.StatusOK, rsp)
}
