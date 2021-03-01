package executor

import "fmt"

// Operator is the interface implemented by all operators.
type Operator interface {
	// Next returns a boolean indicating whether the operator has more work to do.
	Next() bool
	// Execute executes the operation and returns the resulting tuple. Should only
	// be called if Next() returns true.
	Execute() Tuple
}

// Tuple represents a tuple (row) of values.
type Tuple struct {
	Values []Value
}

// Value represents a value and its associated name.
type Value struct {
	Name        string
	StringValue string
}



func NewTuple(inputs ...interface{}) Tuple {
	if len(inputs)%2 != 0 {
		panic(fmt.Sprintf("num inputs must be even, but was: %d", len(inputs)))

	}

	tuple := Tuple{
		Values: make([]Value, 0, len(inputs)/2),
	}

	for i := 0; i < len(inputs); i += 2 {
		tuple.Values = append(tuple.Values, Value{
			Name:        inputs[i].(string),
			StringValue: inputs[i+1].(string),
		})
	}

	return tuple
}
