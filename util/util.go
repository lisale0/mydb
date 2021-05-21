package util

import (
	"errors"
	"strconv"
)

type Expression interface {
	Accept() error
}

type Expr struct {
}

type BinaryExpression struct {
}

func ParseBool(input string) (bool, error) {
	if val, err := strconv.ParseBool(input); err == nil {
		return val, nil
	}
	return false, errors.New("invalid string input to check boolean")
}
