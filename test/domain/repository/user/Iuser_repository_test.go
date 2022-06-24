package user

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/jett/gin-ddd/boot"
	"github.com/magiconair/properties/assert"
	. "github.com/smartystreets/goconvey/convey"

	entity "github.com/jett/gin-ddd/domain/entity/user"
	"github.com/jett/gin-ddd/domain/irepository/user/mock"
	repositoryimpl "github.com/jett/gin-ddd/infrastructure/repository/repositoryimpl/user"
)

// go test -cover ./...
func TestUser_User(t *testing.T) {
	ctrl := gomock.NewController(t) // 初始化 controller
	defer ctrl.Finish()

	userRepo := mock.NewMockIUserRepository(ctrl) // 初始化 mock

	userDao := repositoryimpl.NewUserRepo()
	var ctx context.Context

	Convey("Convey Test Get Userinfo dao", t, func() {
		var err error
		var id uint64

		Convey("Get Userinfo Success", func() {
			err = nil
			id = 1
			repoDataRes := &entity.User{}
			repoDataRes.ID = id
			userRepo.EXPECT().GetUser(ctx, id).AnyTimes().Return(repoDataRes, err)
			dataRes, errRes := userDao.GetUser(ctx, id)
			assert.Equal(t, dataRes.ID, id)
			assert.Equal(t, errRes, err)
		})
	})
}
