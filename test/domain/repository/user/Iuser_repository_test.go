package user

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	. "github.com/smartystreets/goconvey/convey"

	_ "github.com/jett/gin-ddd/boot"
	entity "github.com/jett/gin-ddd/domain/entity/user"
	"github.com/jett/gin-ddd/domain/repository/user/mock"
	dao "github.com/jett/gin-ddd/infrastructure/dao/user"
)

// go test -cover ./...
func TestUser_User(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t) // 初始化 controller
	defer ctrl.Finish()

	userRepo := mock.NewMockIUserRepository(ctrl) // 初始化 mock

	userDao := dao.UserDao{}

	Convey("Convey Test Get Userinfo dao", t, func() {
		var err error
		var id uint64

		Convey("Get Userinfo Success", func() {
			err = nil
			id = 100
			repoDataRes := &entity.User{}
			repoDataRes.ID = id
			userRepo.EXPECT().GetUser(ctx, id).AnyTimes().Return(repoDataRes, err)
			dataRes, errRes := userDao.GetUser(ctx, id)
			assert.Equal(t, dataRes.ID, id)
			assert.Equal(t, errRes, err)
		})
	})
}
