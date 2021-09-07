package lemin

type terrain struct {
	AntsCount  int
	Start, End string
}

func (*terrain) Match() (string, error) {
	return "", nil
}
