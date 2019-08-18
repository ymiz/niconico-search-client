package sort

type Sort struct {
	Field   Field
	OrderBy OrderBy
}

func (s Sort) ToString() string {
	return s.OrderBy.String() + s.Field.String()
}
