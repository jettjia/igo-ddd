package user

import (
	"context"

	entity "github.com/jett/gin-ddd/domain/entity/user"
)

func (this *UserService) GetUser(ctx context.Context, id uint64) (*entity.User, error) {

	return this.userRepo.GetUser(ctx, id)
}
