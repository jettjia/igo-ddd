package user

import (
	"context"

	entity "github.com/jett/gin-ddd/domain/entity/user"
)

func (u *UserService) GetUser(ctx context.Context, id uint64) (*entity.User, error) {

	return u.userRepo.GetUser(ctx, id)
}
