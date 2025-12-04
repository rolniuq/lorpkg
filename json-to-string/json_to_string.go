package jts

import (
	"errors"

	"github.com/bytedance/sonic"
)

func Minify(input string) (string, error) {
	if input == "" {
		return "", errors.New("input is empty")
	}

	var v interface{}

	if err := sonic.Unmarshal([]byte(input), &v); err != nil {
		return "", err
	}

	minified, err := sonic.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(minified), nil
}
