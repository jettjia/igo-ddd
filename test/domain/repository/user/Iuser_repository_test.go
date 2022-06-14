package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	_ "github.com/jett/gin-ddd/boot"
	entity "github.com/jett/gin-ddd/domain/entity/user"
	"github.com/jett/gin-ddd/domain/repository/user/mock"
)

// go test -cover ./...
func Test_User(t *testing.T) {
	ctrl := gomock.NewController(t) // 初始化 controller
	defer ctrl.Finish()

	userRepo := mock.NewMockIUserRepository(ctrl) // 初始化 mock

	//userDao := dao.UserDao{}

	Convey("Convey Test Get Userinfo repository", t, func() {
		var err error
		var id uint64

		Convey("Get Userinfo Log Success", func() {
			err = nil
			id = 100
			repoDataRes := &entity.User{}
			repoDataRes.ID = id
			userRepo.EXPECT().GetUser(id).Return(repoDataRes, err)
			//dataRes, errRes := userDao.GetUser(id)
			//assert.Equal(t, dataRes.ID, id)
			//assert.Equal(t, errRes, err)
		})
	})
}
