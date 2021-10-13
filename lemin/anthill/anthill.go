package anthill

import (
	"errors"
	"fmt"
	"io"
	"sort"
)

func GetAnthillFromLines(lines []string) (*anthill, error) {
	countAnts, err := getCountAntsFromLine(&lines)
	if err != nil {
		return nil, err
	}
	rooms, startRoom, endRoom, err := getRoomsFromLines(&lines)
	if err != nil {
		return nil, err
	}
	err = setPathsFromLines(&lines, rooms)
	if err != nil {
		return nil, err
	}
	return &anthill{AntsCount: countAnts, StepsCount: -1, Rooms: rooms, Start: startRoom, End: endRoom}, nil
}

func (a *anthill) Validate() error {
	if a.AntsCount < 1 {
		return errors.New("invalid number of ants")
	} else if a.Start == "" || a.End == "" {
		return errors.New("no start, end rooms found")
	} else if len(a.Rooms) < 2 {
		return errors.New("invalid number of rooms")
	}
	return nil
}

// Match - To Do
func (a *anthill) Match() error {
	for {
		if !SearchShortPath(a) {
			// path not found, then check for prev path count
			if a.AntsCount > 0 {
				return nil
			} else {
				return errors.New("path not found")
			}
		}
		// checking effective of new short path
		// current steps count < previous steps count
		// if effective then replace result to new (returns true)
		// if not then return previous result (returns false)
		if !CheckEffective(a) {
			return nil
		}
	}
}

func (a *anthill) GenerateResult() {
	sort.Slice(a.Paths, func(i, j int) bool { return a.Paths[i].Len < a.Paths[j].Len })
	steps, antsForEachPath := calcSteps(a.AntsCount, a.Paths)
	// fmt.Printf("MyCalc: %v, Arr: %v\n", steps, antsForEachPath)
	a.Result = make([]string, steps)
	if steps == 1 {
		roomName := a.Paths[0].Front.Room.Name
		for ant := 1; ant <= antsForEachPath[0]; ant++ {
			a.Result[0] += fmt.Sprintf("L%d-%s ", ant, roomName)
		}
	} else {
		curAnt := 1
		for i := range a.Paths {
			start := 0
			for l := a.Paths[i].Front; l != nil; l = l.Next {
				for j := 0; j < antsForEachPath[i]; j++ {
					a.Result[start+j] += fmt.Sprintf("L%d-%s ", curAnt+j, l.Room.Name)
				}
				start++
			}
			curAnt += antsForEachPath[i]
		}
	}
}

// Printing result using Paths on anthill
func (a *anthill) PrintResultByPaths() {
	sort.Slice(a.Paths, func(i, j int) bool { return a.Paths[i].Len < a.Paths[j].Len })
	steps, antsForEachPath := calcSteps(a.AntsCount, a.Paths)
	if steps == 1 {
		roomName := a.Paths[0].Front.Room.Name
		for ant := 1; ant <= antsForEachPath[0]; ant++ {
			fmt.Printf("L%d-%s ", ant, roomName)
		}
		fmt.Println()
	} else {
		fmt.Println("Not Finished!")
	}
}

func (a *anthill) WriteResult(w io.Writer) {
	a.GenerateResult()
	for _, line := range a.Result {
		fmt.Fprintf(w, "%v\n", line)
	}
}
