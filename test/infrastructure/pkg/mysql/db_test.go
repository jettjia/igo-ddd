package mysql

import (
	"fmt"
	"testing"

	_ "github.com/jett/gin-ddd/boot"
	"github.com/jett/gin-ddd/global"
)

type User struct {
	Id       int64  `gorm:"primarykey;type:int;comment:主键" json:"id"`
	Nickname string `gorm:"type:varchar(32); not null; comment昵称;" json:"nickname"`
}

func TestInitDB(t *testing.T) {
	user := User{}
	global.GDB.First(user, 1)

	fmt.Printf("%+v", user)
}
