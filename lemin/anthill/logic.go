package anthill

import (
	"fmt"
	"sort"
)

// search shortest path from start to end with BFS algorithm
// found path state will be saved on UsingOnPath, Parent (needed to check from End Room)
// returns true if found, otherwise false
func SearchShortPath(terrain *anthill) bool {
	usableRoomsList := &list{}
	startRoom := terrain.Rooms[terrain.Start]
	endRoom := terrain.Rooms[terrain.End]
	startRoom.MarkedIn, startRoom.MarkedOut = true, true
	isFind := false
	usedRooms := make(map[*room]bool)
	usedRooms[startRoom] = true
	usableRoomsList.PushBack(startRoom)
	for usableRoomsList.Front != nil {
		current := usableRoomsList.Front.Room
		// mark in and out depends on used on path or not
		if !current.UsingOnPath {
			if addNext(current, endRoom, usedRooms, usableRoomsList, false) {
				isFind = true
				break
			}
		} else {
			if current.MarkedOut {
				if addNext(current, endRoom, usedRooms, usableRoomsList, false) {
					isFind = true
					break
				}
				if current.MarkedIn && addNext(current, endRoom, usedRooms, usableRoomsList, true) {
					isFind = true
					break
				}
			} else {
				if addNext(current, endRoom, usedRooms, usableRoomsList, true) {
					isFind = true
					break
				}
			}
		}
		usableRoomsList.RemoveFront()
	}
	if isFind {
		replaceEdges(startRoom, endRoom)
	}
	clearMarkedInfromation(usedRooms)
	return isFind
}

func CheckEffective(terrain *anthill) bool {
	startRoom, endRoom := terrain.Rooms[terrain.Start], terrain.Rooms[terrain.End]
	i, lenNewPaths := 0, len(startRoom.PathsIn)
	newPaths := make([]*list, lenNewPaths)
	for _, value := range endRoom.PathsIn {
		newPaths[i] = &list{}
		newPaths[i].PushFront(endRoom)
		cur := value
		for cur != startRoom {
			// if len(cur.PathsIn) != 1 {
			// 	log.Fatal(len(cur.PathsIn))
			// }
			newPaths[i].PushFront(cur)
			for _, next := range cur.PathsIn {
				cur = next
			}
		}
		// newPaths[i].PushFront(startRoom)
		i++
	}
	printPaths(newPaths)
	curStepsCount := fastCalcSteps(terrain.AntsCount, newPaths)
	fmt.Printf("%d ants, %d paths, %d steps\n", terrain.AntsCount, len(newPaths), curStepsCount)
	// I Add Start
	sort.Slice(newPaths, func(i, j int) bool { return newPaths[i].Len < newPaths[j].Len })
	tSteps, StepsArr := calcSteps(terrain.AntsCount, newPaths)
	fmt.Printf("MyCalc: %v, Arr: %v\n", tSteps, StepsArr)
	// I Add End
	if terrain.StepsCount == -1 || terrain.StepsCount > curStepsCount {
		terrain.StepsCount = curStepsCount
		terrain.Paths = newPaths
		return true
	}
	return false
}

func (terrain *anthill) generateResult() string {
	return ""
}

func fastCalcSteps(ants int, paths []*list) int {
	steps, lossPerStep := 0, 0
	comingAnts := make(map[int]int)
	for _, value := range paths {
		comingAnts[value.Len]++
	}
	if comingAnts[1] > 0 {
		return 1
	}
	for ants > 0 {
		steps++
		lossPerStep += comingAnts[steps]
		ants -= lossPerStep
	}
	return steps
}

// calcSteps TODO
// PathsLens: [2, 5, 5, 6] | ants: 12
// -=-=-= Start =-=-=-
// i = 0: [ -1,  2,  2,  3] | 12 -1 = 11				| steps: 0+2 = 2

// i = 1: [ -4, -1, -1,  0] | 11 -3 -2 = 6				| steps: 2+3 = 5
// i = 2: [ -5, -2, -2, -1] | 6 -1 -1 -1 -1 = 2			| steps: 5+1 = 6
// OR
// i = 0: [  0,  3,  3,  4] | 12 - 0 = 12				| steps: 0+2 = 2
// i = 1: [ -3,  0,  0,  1] | 12 - 3 = 9				| steps: 2+3 = 5
// i = 2: [ -4, -1, -1,  0] | 9 - 3 = 6					| steps: 5+1 = 6

// -1   : [ -5, -2, -2, -1] | 6 - 4 = 2					| steps: 5+1 = 6
//   DEL: [ -5, -2, -2, -1] | 2/slice.len = 0.5 = 2		| steps: 6+0 = 6
//   MOD: [ -6, -3, -2, -1] | 2%4 = 2, 2-2 = 0			| steps: 6+1 = 7
// On Sending Ants
// [ 1, 1, 1, 1 ] len = 0 | 12    | on st 7 = 0 | will not
// [ x, 1, 1, 1 ] len = 1 | 11    | on st 6
// [ x, 1, 1, 1 ] len = 1 | 10    | on st 5
// [ x, 1, 1, 1 ] len = 1 | 9     | on st 4
// [ x, x, x, 1 ] len = 3 | 6     | on st 3
// [ x, x, x, x ] len = 4 | 2     | on st 2
// [ x, x, -, - ] len = 2 | 0     | on st 1

// Inputs Sorted Paths, and AntsCount should be > 0
func calcSteps(antsCount int, paths []*list) (int, []int) {
	return 0, []int{}
}

func addNext(current, endRoom *room, usedRooms map[*room]bool, usableRoomsList *list, mark bool) bool {
	paths := current.PathsOut
	if mark {
		paths = current.PathsIn
	}
	for _, value := range paths {
		// check if parent doesn't using on path then MarkedIn
		// if both using on path then MarkedOut
		// if value doesn't using on path then usedRooms
		if !current.UsingOnPath && value.UsingOnPath && value.MarkedIn {
			continue
		} else if current.UsingOnPath && value.UsingOnPath && value.MarkedOut {
			continue
		} else if !value.UsingOnPath && usedRooms[value] {
			continue
		}
		// check if room using in path
		if value.UsingOnPath {
			if current.UsingOnPath {
				value.MarkedOut = true
				value.ParentOut = current
			}
			if !value.MarkedIn {
				value.MarkedIn = true
				value.ParentIn = current
			}
		} else {
			value.MarkedIn, value.MarkedOut = true, true
			value.ParentIn, value.ParentOut = current, current
		}
		usedRooms[value] = true
		if value == endRoom {
			return true
		}
		usableRoomsList.PushBack(value)
	}
	return false
}

func replaceEdges(startRoom, endRoom *room) {
	var next *room
	cur, prev := endRoom.ParentIn, endRoom
	if cur == nil {
		cur = endRoom.ParentOut
	}
	delete(endRoom.PathsOut, cur.Name)
	delete(cur.PathsOut, endRoom.Name)
	endRoom.PathsIn[cur.Name] = cur
	for cur != startRoom {
		if _, ok := cur.PathsIn[prev.Name]; ok {
			delete(cur.PathsIn, prev.Name)
			next = cur.ParentIn
		} else {
			next = cur.ParentOut
		}
		if cur.UsingOnPath && next.UsingOnPath {
			cur.PathsOut[next.Name] = next
			next.PathsOut[cur.Name] = cur
		} else {
			delete(cur.PathsOut, next.Name)
			delete(next.PathsOut, cur.Name)
			cur.PathsIn[next.Name] = next
			cur.UsingOnPath = true
		}
		prev, cur = cur, next
	}
	cur.PathsIn[prev.Name] = prev
}

func clearMarkedInfromation(usedRooms map[*room]bool) {
	for key := range usedRooms {
		if len(key.PathsIn) < 1 {
			key.UsingOnPath = false
		}
		key.ParentIn, key.ParentOut = nil, nil
		key.MarkedIn, key.MarkedOut = false, false
	}
}
