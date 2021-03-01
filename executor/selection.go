package executor

type SelectionOperator struct {
	selection map[string]bool
	child Operator
	idx int
}


func NewSelectionOperator (selection map[string]bool, child Operator) Operator {
	return &SelectionOperator{
		selection: selection,
		child: child,
		idx: -1,
	}
}

func (s *SelectionOperator) Next() bool {
	s.idx += 1
	if s._hasValidNext(){
		return true
	}
	return false
}
func (s *SelectionOperator) Execute() Tuple {
	tuple := s.child.Execute()
	filter := tuple.Values[:0]
	for _, value := range tuple.Values {
		if _, ok := s.selection[value.Name]; ok {
			newValue := Value{
			value.Name,
			value.StringValue,
			}
			filter = append(filter, newValue)
		}
	}
	tuple.Values = filter
	return tuple
}
func (s *SelectionOperator) _hasValidNext() bool {
	return s.child.Next()

}