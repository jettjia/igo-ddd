package aggregate

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	. "github.com/smartystreets/goconvey/convey"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	repoMock "github.com/jettjia/go-ddd-demo/domain/irepository/sys/mock"
	"github.com/jettjia/go-ddd-demo/types"
)

func Test_CreateSysMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		ctx       context.Context
		err       error
		id        uint64
		sysMenuEn entity.SysMenu
	)

	repoMockIns := repoMock.NewMockISysMenuRepo(ctrl)

	ins := &sysMenuAgg{
		sysMenuImpl: repoMockIns,
	}

	Convey("Convey Test CreateSysMenu", t, func() {
		Convey("CreateSysMenu Success", func() {
			err = nil

			repoMockIns.EXPECT().Create(ctx, gomock.Any()).AnyTimes().Return(id, err)

			idRsp, errRsp := ins.CreateSysMenu(ctx, &sysMenuEn)
			assert.Equal(t, idRsp, id)
			assert.Equal(t, errRsp, err)
		})
	})
}

func Test_DeleteSysMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		ctx       context.Context
		err       error
		sysMenuEn entity.SysMenu
	)

	repoMockIns := repoMock.NewMockISysMenuRepo(ctrl)

	ins := &sysMenuAgg{
		sysMenuImpl: repoMockIns,
	}

	Convey("Convey Test DeleteSysMenu", t, func() {
		Convey("DeleteSysMenu Success", func() {
			err = nil

			repoMockIns.EXPECT().Delete(ctx, gomock.Any()).AnyTimes().Return(err)

			errRsp := ins.DeleteSysMenu(ctx, &sysMenuEn)
			assert.Equal(t, errRsp, err)
		})
	})
}

func Test_UpdateSysMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		ctx       context.Context
		err       error
		sysMenuEn entity.SysMenu
	)

	repoMockIns := repoMock.NewMockISysMenuRepo(ctrl)

	ins := &sysMenuAgg{
		sysMenuImpl: repoMockIns,
	}

	Convey("Convey Test UpdateSysMenu", t, func() {
		Convey("UpdateSysMenu Success", func() {
			err = nil

			repoMockIns.EXPECT().Update(ctx, gomock.Any()).AnyTimes().Return(err)

			errRsp := ins.UpdateSysMenu(ctx, &sysMenuEn)
			assert.Equal(t, errRsp, err)
		})
	})
}

func Test_FindSysMenuById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		ctx       context.Context
		err       error
		sysMenuEn entity.SysMenu
		id        uint64
	)

	repoMockIns := repoMock.NewMockISysMenuRepo(ctrl)

	ins := &sysMenuAgg{
		sysMenuImpl: repoMockIns,
	}

	Convey("Convey Test FindSysMenuById", t, func() {
		Convey("FindSysMenuById Success", func() {
			err = nil

			repoMockIns.EXPECT().FindById(ctx, gomock.Any()).AnyTimes().Return(&sysMenuEn, err)

			sysMenuEnRsp, errRsp := ins.FindSysMenuById(ctx, id)
			assert.Equal(t, errRsp, err)
			assert.Equal(t, sysMenuEnRsp.Id, id)
		})
	})
}

func Test_FindSysMenuByQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		ctx       context.Context
		err       error
		sysMenuEn entity.SysMenu
		id        uint64
		query     []*types.Query
	)

	repoMockIns := repoMock.NewMockISysMenuRepo(ctrl)

	ins := &sysMenuAgg{
		sysMenuImpl: repoMockIns,
	}

	Convey("Convey Test FindSysMenuByQuery", t, func() {
		Convey("FindSysMenuByQuery Success", func() {
			err = nil

			repoMockIns.EXPECT().FindByQuery(ctx, gomock.Any()).AnyTimes().Return(&sysMenuEn, err)

			sysMenuEnRsp, errRsp := ins.FindSysMenuByQuery(ctx, query)
			assert.Equal(t, errRsp, err)
			assert.Equal(t, sysMenuEnRsp.Id, id)
		})
	})
}

func Test_FindSysMenuAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		ctx         context.Context
		err         error
		sysMenuList []*entity.SysMenu
		query       []*types.Query
	)

	repoMockIns := repoMock.NewMockISysMenuRepo(ctrl)

	ins := &sysMenuAgg{
		sysMenuImpl: repoMockIns,
	}

	Convey("Convey Test FindSysMenuAll", t, func() {
		Convey("FindSysMenuAll Success", func() {
			err = nil

			repoMockIns.EXPECT().FindAll(ctx, gomock.Any()).AnyTimes().Return(sysMenuList, err)

			_, errRsp := ins.FindSysMenuAll(ctx, query)
			assert.Equal(t, errRsp, err)
		})
	})
}

func Test_FindSysMenuPage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		ctx         context.Context
		err         error
		sysMenuList []*entity.SysMenu
		query       []*types.Query
		reqPage     *types.PageData
		reqSort     *types.SortData
		resPage     *types.PageData
	)

	repoMockIns := repoMock.NewMockISysMenuRepo(ctrl)

	ins := &sysMenuAgg{
		sysMenuImpl: repoMockIns,
	}

	Convey("Convey Test FindSysMenuPage", t, func() {
		Convey("FindSysMenuPage Success", func() {
			err = nil

			repoMockIns.EXPECT().FindPage(ctx, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(sysMenuList, resPage, err)

			_, _, errRsp := ins.FindSysMenuPage(ctx, query, reqPage, reqSort)
			assert.Equal(t, errRsp, err)
		})
	})
}
