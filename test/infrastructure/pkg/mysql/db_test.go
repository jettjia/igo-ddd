package mysql

import (
	"testing"

	_ "github.com/jett/gin-ddd/boot"
	"github.com/jett/gin-ddd/global"
)

func TestInitDB(t *testing.T) {
	err := global.GDB.DB().Ping()
	if err != nil {
		t.Fatal("mysql conn err :", err)
	}
}
