package util

import "github.com/oklog/ulid/v2"

func Ulid() string {
	return ulid.Make().String()
}
