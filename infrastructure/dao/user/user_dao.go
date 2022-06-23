package dao

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	entity "github.com/jett/gin-ddd/domain/entity/user"
	userRepository "github.com/jett/gin-ddd/domain/repository/user"
	"github.com/jett/gin-ddd/global"
	"github.com/jett/gin-ddd/infrastructure/consts"
)

var _ userRepository.IUserRepository = (*UserDao)(nil)

type UserDao struct {
	db *gorm.DB
}

func NewUserRepo() *UserDao {
	return &UserDao{
		db: global.GDB,
	}
}

func (this *UserDao) getCacheKey(data string) string {
	return fmt.Sprintf("%s%s", consts.UserCacheKey, data)
}

func (this *UserDao) SaveUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := this.db.Create(&user).Error
	if err != nil {
		global.GLog.Errorln(err.Error())
		return nil, err
	}

	return user, nil
}

func (this *UserDao) GetUser(ctx context.Context, id uint64) (*entity.User, error) {
	var (
		user entity.User
	)
	err := this.db.First(&user, id).Error
	if err != nil {
		global.GLog.Errorln(err.Error())
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("food not found")
	}
	return &user, nil
}

func (this *UserDao) GetUserByName(ctx context.Context, nickname string) (*entity.User, error) {
	var (
		user entity.User
	)
	err := this.db.Where("nickname", nickname).Error
	if err != nil {
		global.GLog.Errorln(err.Error())
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("food not found")
	}
	return &user, nil
}
