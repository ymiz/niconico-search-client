package filters

type Filter struct {
	Field     Field
	Condition string
	Value     string
}

func (f Filter) GenerateKeyString() string {
	return "filters[" + f.Field.String() + "][" + f.Condition + "]"
}
