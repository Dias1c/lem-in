package anthill

import (
	"fmt"
	"io"
	"sort"
)

// pathsOfListToSlice - convertiong list of rooms to slice of rooms
func pathsOfListToSlice(paths []*list) [][]*room {
	result := make([][]*room, len(paths))
	for i, path := range paths {
		result[i] = make([]*room, path.Len)
		j := 0
		for node := path.Front; node != nil; node = node.Next {
			result[i][j] = node.Room
			j++
		}
	}
	return result
}

// WriteResult - write result with writer
func (r *Result) WriteResult(w io.Writer) {
	sort.Slice(r.Paths, func(i, j int) bool { return r.Paths[i].Len < r.Paths[j].Len })
	steps, antsForEachPath := calcSteps(r.AntsCount, r.Paths)
	if steps == 1 {
		roomName := r.Paths[0].Front.Room.Name
		for ant := 1; ant <= antsForEachPath[0]; ant++ {
			fmt.Fprintf(w, "L%d-%s ", ant, roomName)
		}
		fmt.Fprintln(w)
	} else {
		paths := pathsOfListToSlice(r.Paths)
		a, b := &antQueue{}, &antQueue{}
		antNum := 1
		for i := 1; i <= steps; i++ {
			cur := a.Dequeue()
			for cur != nil {
				// fmt.Fprintf(output, "L%d-%s ", cur.Num, result.Paths[cur.Path][cur.Pos].Name)
				w.Write([]byte(fmt.Sprintf("L%d-%s ", cur.Num, paths[cur.Path][cur.Pos].Name)))
				cur.Pos++
				if cur.Pos < len(paths[cur.Path]) {
					b.EnqueueAnt(cur)
				}
				cur = a.Dequeue()
			}
			for j, v := range antsForEachPath {
				if v > 0 {
					// fmt.Fprintf(output, "L%d-%s ", antNum, result.Paths[j][0].Name)
					w.Write([]byte(fmt.Sprintf("L%d-%s ", antNum, paths[j][0].Name)))
					antsForEachPath[j]--
					b.Enqueue(antNum, j, 1)
					antNum++
				}
			}
			t := a
			a = b
			b = t
			fmt.Fprintln(w)
		}
	}
}
