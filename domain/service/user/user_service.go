package user

import (
	"context"

	entity "github.com/jett/gin-ddd/domain/entity/user"
	repository "github.com/jett/gin-ddd/domain/irepository/user"
	dao "github.com/jett/gin-ddd/infrastructure/repository/repositoryimpl/user"
)

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: dao.NewUserRepo(),
	}
}

func (u *UserService) GetUser(ctx context.Context, id uint64) (*entity.User, error) {

	return u.userRepo.GetUser(ctx, id)
}
