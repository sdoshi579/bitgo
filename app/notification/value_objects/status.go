package value_objects

type Status int

const (
	StatusOutstanding Status = iota
	StatusSent
	StatusFailed
)

var AllStatusString = []string{"Outstanding", "Sent", "Failed"}

func (s Status) string() string {
	return AllStatusString[s]
}

func GetStatusEnum(status string) Status {
	for i, s := range AllStatusString {
		if s == status {
			return Status(i)
		}
	}
	return StatusOutstanding
}
