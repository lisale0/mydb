package executor

type ScanOperator struct {
	tuples []Tuple
	idx    int
}

func NewScanOperator(tuples []Tuple) Operator {
	return &ScanOperator{
		tuples: tuples,
		idx:    -1,
	}
}

func (s *ScanOperator) Next() bool {
	s.idx += 1
	return s.idx < len(s.tuples)
}

func (s *ScanOperator) Execute() Tuple {
	return s.tuples[s.idx]
}
