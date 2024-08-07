package utils

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func AssertEquals[T comparable](a, b T) {
	if a != b {
		panic(errors.New("not equals"))
	}
}

func Neat(v interface{}) string {
	data, err := NeatBytes(v)
	if err != nil {
		return "" //when the result is empty string, means wrong
	}
	return string(data)
}

func NeatBytes(v interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, errors.WithMessage(err, "marshal object is wrong")
	}
	return data, nil
}

// PtrX 给任何值取地址，特别是当参数为数字0，1，2，3或者字符串"a", "b", "c"的时候
func PtrX[T any](v T) *T {
	return &v
}

// VOr0 给任何地址取值，当是空地址时返回 zero 即类型默认的零值
func VOr0[T any](v *T) T {
	if v != nil {
		return *v
	} else {
		var zero T
		return zero
	}
}
