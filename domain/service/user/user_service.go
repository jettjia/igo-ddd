package user

import (
	"context"
	entity "github.com/jett/gin-ddd/domain/entity/user"
	dao "github.com/jett/gin-ddd/infrastructure/repository/repositoryimpl/user"

	repository "github.com/jett/gin-ddd/domain/irepository/user"
)

// UserService user service interface
type UserService interface {
	FindUserByID(ctx context.Context, id uint64) (*entity.User, error)
}

var _ UserService = (*userService)(nil)

type userService struct {
	userRepo repository.IUserRepository
}

func (u *userService) FindUserByID(ctx context.Context, id uint64) (*entity.User, error) {
	return u.userRepo.GetUser(ctx, id)
}

func NewUserService() *userService {
	return &userService{
		userRepo: dao.NewUserRepo(),
	}
}
