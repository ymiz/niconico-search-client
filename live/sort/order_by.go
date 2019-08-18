package sort

type OrderBy int

func (o OrderBy) String() string {
	switch o {
	case Asc:
		return "+"
	case Desc:
		return "-"
	default:
		return "Unknown"
	}
}

const (
	OrderUnknown OrderBy = iota
	Asc
	Desc
)
