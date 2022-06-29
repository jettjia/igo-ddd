package user

import (
	"context"
	entity "github.com/jettjia/gin-ddd/domain/entity/user"
	repository "github.com/jettjia/gin-ddd/domain/irepository/user"
	dao "github.com/jettjia/gin-ddd/infrastructure/repository/repositoryimpl/user"
)

// UserService user service interface
type UserService interface {
	FindUserByID(ctx context.Context, id uint64) (*entity.User, error)
}

var _ UserService = (*userService)(nil)

type userService struct {
	userRepo repository.IUserRepository
}

func NewUserService() *userService {
	return &userService{
		userRepo: dao.NewUserRepo(),
	}
}

func (u *userService) FindUserByID(ctx context.Context, id uint64) (*entity.User, error) {
	return u.userRepo.GetUser(ctx, id)
}
