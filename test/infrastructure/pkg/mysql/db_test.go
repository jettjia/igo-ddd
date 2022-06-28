package mysql

import (
	"testing"

	"github.com/jett/gin-ddd/global"
	_ "github.com/jett/gin-ddd/test"
)

func TestInitDB(t *testing.T) {
	err := global.GDB.DB().Ping()
	if err != nil {
		t.Fatal("mysql conn err :", err)
	}
}
