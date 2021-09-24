package anthill

import "container/list"

type paths struct {
	Name  string
	Paths []*list.List
}

// TODO
func (p *paths) CountSteps(countAnts int) int {
	return 0
}
