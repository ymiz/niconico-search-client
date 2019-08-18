package fields

import "strings"

type Fields struct {
	Field []Field
}

func (f Fields) HasField() bool {
	return len(f.Field) > 0 && f.Field[0] != Unknown
}

func (f Fields) ToString() string {
	var fieldStrings []string
	for _, field := range f.Field {
		fieldStrings = append(fieldStrings, field.String())
	}
	return strings.Join(fieldStrings, ",")
}
