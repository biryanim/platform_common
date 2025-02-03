package filter

type Condition struct {
	Key   string
	Value interface{}
}

type Filter struct {
	Conditions []Condition
}

func New(conditions ...Condition) *Filter {
	return &Filter{
		Conditions: conditions,
	}
}

func (f *Filter) AddCondition(condition Condition) {
	f.Conditions = append(f.Conditions, condition)
}
