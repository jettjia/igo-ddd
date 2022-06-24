package converter

import (
	entity "github.com/jett/gin-ddd/domain/entity/user"
	po "github.com/jett/gin-ddd/infrastructure/repository/po/user"
)

type UserConv struct {
}

func E2PUser(user *entity.User) *po.User {
	var poUser po.User

	poUser.ID = user.ID
	poUser.Passport = user.Passport
	poUser.Password = user.Password
	poUser.Nickname = user.Nickname
	poUser.CreatedAt = user.CreatedAt
	poUser.UpdatedAt = user.UpdatedAt
	poUser.DeletedAt = user.DeletedAt

	return &poUser
}

func P2EUser(user *po.User) *entity.User {
	var enUser entity.User

	enUser.ID = user.ID
	enUser.Passport = user.Passport
	enUser.Password = user.Password
	enUser.Nickname = user.Nickname
	enUser.CreatedAt = user.CreatedAt
	enUser.UpdatedAt = user.UpdatedAt
	enUser.DeletedAt = user.DeletedAt

	return &enUser
}
