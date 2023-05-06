package memory

type Memory interface {
}

type memory struct{}

func New() Memory {
	return &memory{}
}
