package executor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLimitOperatorLessTuplesThanLimit(t *testing.T) {
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
	scanOp := NewScanOperator(tuples)
	limitOp := NewLimitOperator(len(tuples)+1, scanOp)

	// All tuples should be returned.
	for _, tuple := range tuples {
		is.Equal(true, limitOp.Next())
		is.Equal(tuple, limitOp.Execute())
	}
	is.Equal(false, limitOp.Next())
}

func TestLimitOperatorMoreTuplesThanLimit(t *testing.T) {
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
	scanOp := NewScanOperator(tuples)
	limitOp := NewLimitOperator(len(tuples)-1, scanOp)

	for _, tuple := range tuples[:len(tuples)-1] {
		is.Equal(true, limitOp.Next())
		is.Equal(tuple, limitOp.Execute())
	}
	is.Equal(false, limitOp.Next())
}
