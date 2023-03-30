package util

import (
	"fmt"
	"testing"
)

func Test_Ulid(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Ulid() + "\n")
	}
}
