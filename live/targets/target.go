package targets

import "strings"

type Targets struct {
	Field []Field
}

func (t Targets) ToString() string {
	var fieldStrings []string
	for _, field := range t.Field {
		fieldStrings = append(fieldStrings, field.String())
	}
	return strings.Join(fieldStrings, ",")
}
