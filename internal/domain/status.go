package domain

type Status int

const (
	Unknown Status = iota
	Unselected
	Selected
	Gifted
)

func NewStatus(s string) Status {
	switch s {
	case "unselected":
		return Unselected
	case "selected":
		return Selected
	case "gifted":
		return Gifted
	default:
		return Unknown
	}
}

func (s Status) String() string {
	switch s {
	case Unselected:
		return "unselected"
	case Selected:
		return "selected"
	case Gifted:
		return "gifted"
	default:
		return "unknown"
	}
}
