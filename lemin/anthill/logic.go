package anthill

import "fmt"

// SearchShortPath - search shortest path from start to end with Bellman-Ford algorithm (Suurballe`s algorithm).
// found path state will be saved on UsingOnPath, Parent (needed to check from End Room)
// returns true if found, otherwise false
func searchShortPath(terrain *anthill) bool {
	usableRoomsQueue := &sortedQueue{}
	visitedRooms := make(map[*room]bool)
	startRoom := terrain.Rooms[terrain.Start]
	endRoom := terrain.Rooms[terrain.End]
	visitedRooms[startRoom] = true
	usableRoomsQueue.Enqueue(startRoom)
	for usableRoomsQueue.Front != nil && !visitedRooms[endRoom] {
		current := usableRoomsQueue.Dequeue().Room
		// fmt.Printf("Current: %s\n", current.Name)
		for next, value := range current.Paths {
			if value == BLOCKED || (current.Separated && !current.Marked && value == STABLE) {
				continue
			}
			addNext(current, next, value, visitedRooms, usableRoomsQueue)
		}
		// fmt.Print("Queue state: ")
		// usableRoomsQueue.DebugPrint()
		// fmt.Println()
	}
	isFind := visitedRooms[endRoom]
	// debugVisitedRooms(visitedRooms)
	if isFind {
		// to do replace edges
		replaceEdges(startRoom, endRoom)
	}
	startRoom.Separated = false
	endRoom.Separated = false
	return isFind
}

func debugVisitedRooms(visitedRooms map[*room]bool) {
	fmt.Println("### Debug Visited Rooms ###")
	i := 1
	for key := range visitedRooms {
		parentName := "nil"
		if key.Parent != nil {
			parentName = key.Parent.Name
		}
		fmt.Printf("%d. Room: %s\n\tParent: %s\n\tWeight: %d\n", i, key.Name, parentName, key.Weight)
		i++
	}
	fmt.Println("###########################\n")
}

// addNext - add into usableRoomsQueue next room
func addNext(cur, next *room, state int, visitedRooms map[*room]bool, usableRoomsQueue *sortedQueue) {
	weight := cur.Weight + cur.Paths[next]
	if !visitedRooms[next] {
		markNext(cur, next, weight, visitedRooms)
		usableRoomsQueue.Enqueue(next)
		return
	}
	if (next.Separated && !next.Marked) && cur.Paths[next] == RESERVED {
		if next.Weight >= weight {
			markNext(cur, next, weight, visitedRooms)
		} else {
			next.Marked = true
		}
	} else if next.Weight > weight {
		markNext(cur, next, weight, visitedRooms)
	} else {
		// if next.Separated {
		// 	fmt.Printf("Returned: %s|sep\n", next.Name)
		// } else {
		// 	fmt.Printf("Returned: %s\n", next.Name)
		// }
		return
	}
	usableRoomsQueue.SortEnqueue(next)
}

// markNext - mark flags and set weight
func markNext(parent, cur *room, weight int, visitedRooms map[*room]bool) {
	// if !visitedRooms[cur] {
	// 	fmt.Printf("#Visited first: %s\n", cur.Name)
	// }
	// if cur.Separated {
	// 	fmt.Printf("Added: %s|sep\t%s\t%d\n", cur.Name, parent.Name, weight)
	// } else {
	// 	fmt.Printf("Added: %s\t%s\t%d\n", cur.Name, parent.Name, weight)
	// }
	visitedRooms[cur] = true
	cur.Weight = weight
	cur.Parent = parent
	if !cur.Separated {
		return
	}
	if parent.Paths[cur] == RESERVED {
		cur.Marked = true
	} else {
		cur.Marked = false
	}
}

// replaceEdges - replace edges for finded paths. (Suurballe`s algorithm)
func replaceEdges(startRoom, endRoom *room) {
	r := endRoom
	for r != startRoom {
		if r.Paths[r.Parent] == STABLE {
			r.Parent.Separated = true
			r.Separated = true
			r.Paths[r.Parent] = RESERVED
			r.Parent.Paths[r] = BLOCKED
		} else {
			r.Parent.Separated = false
			r.Paths[r.Parent] = STABLE
			r.Parent.Paths[r] = STABLE
		}
		r = r.Parent
	}
}

// checking effective of new short path
// current steps count < previous steps count
// if effective then replace result to new (returns true)
// if not then return previous result (returns false)
func checkEffective(terrain *anthill) bool {
	startRoom, endRoom := terrain.Rooms[terrain.Start], terrain.Rooms[terrain.End]
	i, lenNewPaths := 0, 0
	for _, value := range startRoom.Paths {
		if value == BLOCKED {
			lenNewPaths++
		}
	}
	newPaths := make([]*list, lenNewPaths)
	for key, value := range startRoom.Paths {
		if value == BLOCKED {
			newPaths[i] = &list{}
			cur := key
			for cur != endRoom {
				newPaths[i].PushBack(cur)
				for next, vNext := range cur.Paths {
					if vNext == BLOCKED {
						cur = next
						break
					}
				}
			}
			newPaths[i].PushBack(endRoom)
			i++
		}
	}
	// fmt.Println("### Debug Paths ###")
	// for _, v := range newPaths {
	// 	fr := v.Front
	// 	for fr != v.Back {
	// 		fmt.Printf("%s -> ", fr.Room.Name)
	// 		fr = fr.Next
	// 	}
	// 	fmt.Printf("%s\n", fr.Room.Name)
	// }
	// fmt.Println("###################\n")
	curStepsCount := fastCalcSteps(terrain.AntsCount, newPaths)
	// fmt.Printf("%d ants, %d paths, %d steps\n", terrain.AntsCount, len(newPaths), curStepsCount)
	if terrain.StepsCount == 0 || terrain.StepsCount > curStepsCount {
		terrain.StepsCount = curStepsCount
		terrain.Result.Paths = newPaths
		return curStepsCount != 1
	}
	return false
}

// fastCalcSteps - calculate steps for paths and ants count
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

// calcSteps LOGIC
// Example: PathsLens: [2, 5, 5, 6] | ants: 12
// -=-=-= Start =-=-=-
// Steps = 6 = slice[slice.len-1] | (lastElem + 1) - eachElem -> [5, 2, 2, 1]
// ants = ants - sumElems
//   DEL: [5, 2, 2, 1] | ants/slice.Len = 2/slice.len = 0.5 | ants = 2		| steps: 6+0 = 6
//   MOD: [6, 3, 2, 1] | ants%slice.Len = 2%4 = 2, 2-2 = 0	| ants = 0		| steps: 6+1 = 7
// On Sending Ants
// [ 1, 1, 1, 1 ] len = 0 | 12    | on st 7 = 0 | will not
// [ x, 1, 1, 1 ] len = 1 | 11    | on st 6
// [ x, 1, 1, 1 ] len = 1 | 10    | on st 5
// [ x, 1, 1, 1 ] len = 1 | 9     | on st 4
// [ x, x, x, 1 ] len = 3 | 6     | on st 3
// [ x, x, x, x ] len = 4 | 2     | on st 2
// [ x, x, -, - ] len = 2 | 0     | on st 1

// Inputs Sorted Paths, and AntsCount should be > 0
// Function designed for the optimal number of paths for ants count
func calcSteps(antsCount int, sortedPaths []*list) (int, []int) {
	if len(sortedPaths) < 1 {
		return 0, []int{}
	}
	if sortedPaths[0].Len == 1 {
		return 1, []int{antsCount}
	}
	// Create Result
	lenPaths := len(sortedPaths)
	result := make([]int, lenPaths)
	steps, lastElem := sortedPaths[lenPaths-1].Len, sortedPaths[lenPaths-1].Len+1
	for i := 0; i < lenPaths; i++ {
		result[i] = lastElem - sortedPaths[i].Len
		antsCount -= result[i]
	}
	if antsCount > 0 {
		if antsCount >= lenPaths {
			del := antsCount / lenPaths
			antsCount %= lenPaths
			steps += del
			for i := 0; i < lenPaths; i++ {
				result[i] += del
			}
		}
		if antsCount > 0 {
			steps++
			for i := 0; i < antsCount; i++ {
				result[i]++
			}
			antsCount = 0
		}
	}
	return steps, result
}
