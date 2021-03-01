package executor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanWithTuples(t *testing.T) {
	is := assert.New(t)
	tuples := []Tuple{
		NewTuple(
			"id", "student1",
			"gender", "male"),
		NewTuple(
			"id", "student2",
			"gender", "female"),
		NewTuple(
			"id", "student3",
			"gender", "female"),
	}
	op := NewScanOperator(tuples)
	for _, tuple := range tuples {
		is.Equal(true, op.Next())
		is.Equal(tuple, op.Execute())
	}
	is.Equal(false, op.Next())

}
