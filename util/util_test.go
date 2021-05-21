package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseBoolStrTrue(t *testing.T) {
	is := assert.New(t)
	input := "true"
	output, _ := ParseBool(input)
	is.Equal(true, output)

}

func TestParseBoolStrFalse(t *testing.T) {
	is := assert.New(t)
	input := "false"
	output, _ := ParseBool(input)
	is.Equal(false, output)

}
