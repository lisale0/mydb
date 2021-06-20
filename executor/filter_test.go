package executor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterOnTrueExpr(t *testing.T) {
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
	expr := "true"
	filterOp := NewFilterOperator(tuples, expr, scanOp)

	//expected to return all tuples provided
	for _, tuple := range tuples {
		is.Equal(true, filterOp.Next())
		is.Equal(tuple, filterOp.Execute())
	}
}

//func TestFilterSimpleBinaryExpr(t *testing.T) {
//	is := assert.New(t)
//	tuples := []Tuple{
//		NewTuple(
//			"id", "student1",
//			"gender", "male"),
//		NewTuple(
//			"id", "student2",
//			"gender", "female"),
//	}
//	scanOp := NewScanOperator(tuples)
//	expr := "id == male"
//	filterOp := NewFilterOperator(tuples, expr, scanOp)
//
//	//expected to return all tuples provided
//	is.Equal(true, filterOp.Next())
//	is.Equal(tuples[0], filterOp.Execute())
//
//	is.Equal(false, filterOp.Next())
//	is.Equal(tuples[1], filterOp.Execute())
//
//}
