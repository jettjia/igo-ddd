package repository

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	entity "github.com/jett/gin-ddd/domain/entity/user"
	"github.com/jett/gin-ddd/global"
	"github.com/jett/gin-ddd/infrastructure/consts"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (this *UserRepo) getCacheKey(data string) string {
	return fmt.Sprintf("%s%s", consts.UserCacheKey, data)
}

func (this *UserRepo) SaveUser(user *entity.User) (*entity.User, error) {
	err := global.GDB.Create(&user).Error
	if err != nil {
		global.GLog.Errorln(err.Error())
		return nil, err
	}

	return user, nil
}

func (this *UserRepo) GetUser(id uint64) (*entity.User, error) {
	var (
		user entity.User
	)
	err := global.GDB.First(&user, id).Error
	if err != nil {
		global.GLog.Errorln(err.Error())
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("food not found")
	}
	return &user, nil
}

func (this *UserRepo) GetUserByName(nickname string) (*entity.User, error) {
	var (
		user entity.User
	)
	err := global.GDB.Where("nickname", nickname).Error
	if err != nil {
		global.GLog.Errorln(err.Error())
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("food not found")
	}
	return &user, nil
}
