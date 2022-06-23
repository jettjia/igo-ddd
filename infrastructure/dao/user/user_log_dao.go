package dao

import (
	"context"
	"fmt"

	entity "github.com/jett/gin-ddd/domain/entity/user"
	userRepository "github.com/jett/gin-ddd/domain/repository/user"
	"github.com/jett/gin-ddd/global"
	"github.com/jett/gin-ddd/infrastructure/consts"
)

var _ userRepository.IUserLogRepository = (*UserLog)(nil)

type UserLog struct {
}

func (a *UserLog) getCacheKey(data string) string {
	return fmt.Sprintf("%s%s", consts.UserLogCacheKey, data)
}

func (a *UserLog) SaveLog(ctx context.Context, log *entity.UserLog) (*entity.UserLog, error) {
	err := global.GDB.Create(&log).Error
	if err != nil {
		global.GLog.Errorln(err.Error())
		return nil, err
	}

	return log, nil
}
