package user

import entity "github.com/jett/gin-ddd/domain/entity/user"

func (this *UserService) GetUser(id uint64) (*entity.User, error) {

	return this.userRepo.GetUser(id)
}
