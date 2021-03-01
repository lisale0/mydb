package executor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionWithScan(t *testing.T) {
	is := assert.New(t)
	tuples := []Tuple{
		NewTuple(
			"id", "student1",
			"gender", "male",
			"age", "21"),
	}
	values := []Value{
		Value{
			"gender",
			"male",
		},
		Value{
			"age",
			"21",
		},
	}

	expected := Tuple{values}
	scanOp := NewScanOperator(tuples)
	selection := map[string]bool{"gender": true, "age": true}
	selectionOp := NewSelectionOperator(selection, scanOp)
	is.Equal(true, selectionOp.Next())
	is.Equal(expected, selectionOp.Execute())
}
