package repository

import (
	"context"

	entity "github.com/jett/gin-ddd/domain/entity/user"
)

//go:generate mockgen --source ./Iuser_repository.go --destination ./mock/mock_user.go --package mock
type IUserRepository interface {
	SaveUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id uint64) (*entity.User, error)
	GetUserByName(ctx context.Context, nickname string) (*entity.User, error)
}
