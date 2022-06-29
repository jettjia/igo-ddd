package irepository

import (
	"context"

	entity "github.com/jettjia/go-ddd/domain/entity/user"
)

//go:generate mockgen --source ./Iuser_log_repository.go --destination ./mock/mock_user_log.go --package mock
type IUserLogRepository interface {
	SaveLog(ctx context.Context, log *entity.UserLog) (*entity.UserLog, error)
}
