package executor

import (
	"errors"
	"github.com/lisale0/mydb/util"
	"strings"
)

type FilterOperator struct {
	tuples []Tuple
	expr   string
	idx    int
	child  Operator
}

func NewFilterOperator(tuples []Tuple, expr string, child Operator) Operator {
	return &FilterOperator{
		tuples: tuples,
		expr:   expr,
		idx:    -1,
		child:  child,
	}
}

func (f *FilterOperator) Next() bool {
	f.idx += 1
	if f._isValidTuple() && f._hasValidNext() {
		f.child.Execute()
		return true
	}
	return false
}

func (f *FilterOperator) Execute() Tuple {
	return f.tuples[f.idx]
}

// check of the tuple is valid based on a expression evaluated
func (f *FilterOperator) _isValidTuple() bool {
	return f._evaluateExpression()
}

func (f *FilterOperator) _evaluateExpression() bool {
	if _, err := util.ParseBool(f.expr); err == nil {
		return true
	}
	if _, err := f._evaluateBinary(); err == nil {
		return true
	}

	return false
}

func (f *FilterOperator) _evaluateBinary() (bool, error) {
	s := strings.Split(f.expr, " ")
	if len(s) < 3 {
		return false, errors.New("invalid binary expression")
	}
	left := s[0]
	op := s[1]
	right := s[2]

	if op == "==" {
		if f._checkEquality(left, right) {
			return true, nil
		}
	}
	return false, errors.New("not a match")
}

func (f *FilterOperator) _checkEquality(left, right string) bool {
	tupleLeftVal := f.tuples[f.idx].Values[0].Name
	tupleRightVal := f.tuples[f.idx].Values[1].StringValue
	return (tupleLeftVal == left) && tupleRightVal == right
}
func (f *FilterOperator) _hasValidNext() bool {
	return f.child.Next()
}
