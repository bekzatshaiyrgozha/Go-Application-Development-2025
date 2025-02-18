package practice1_unpack_string

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(srcStr string) (string, error) {
	// write code here
	return srcStr, ErrInvalidString
}
