package mysql

import (
	"testing"

	"github.com/jettjia/go-ddd/global"
	_ "github.com/jettjia/go-ddd/test"
)

func TestInitDB(t *testing.T) {
	err := global.GDB.DB().Ping()
	if err != nil {
		t.Fatal("mysql conn err :", err)
	}
}
