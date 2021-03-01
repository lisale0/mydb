package executor

type LimitOperator struct {
	tupleLimit int
	idx        int
	child      Operator
}


func NewLimitOperator(tupleLen int, child Operator) Operator {
	return &LimitOperator{
		tupleLimit: tupleLen,
		child: child,
		idx: 0,
	}
}


func (l *LimitOperator) Next() bool {
	valid := l._withinLimit() && l._hasValidNextChild()
	if valid {
		l.idx += 1
	}
	return valid
}

func (l *LimitOperator) Execute() Tuple {
	return l.child.Execute()
}

func (l *LimitOperator) _withinLimit() bool {
	return l.idx < l.tupleLimit
}

func (l *LimitOperator) _hasValidNextChild() bool {
	return l.child.Next()
}

