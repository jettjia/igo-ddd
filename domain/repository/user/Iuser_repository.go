package repository

import entity "github.com/jett/gin-ddd/domain/entity/user"

//go:generate mockgen --source ./Iuser_repository.go --destination ./mock/mock_user.go --package mock
type IUserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUserByName(nickname string) (*entity.User, error)
}
