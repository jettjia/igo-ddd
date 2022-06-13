package repository

import entity "github.com/jett/gin-ddd/domain/entity/user"

//go:generate mockgen --source ./user_log.go --destination ./mock/mock_user_log.go --package mock
type IUserLogRepository interface {
	SaveLog(log *entity.UserLog) (*entity.UserLog, error)
}
