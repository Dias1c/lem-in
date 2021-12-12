package anthill

// SearchShortPath - search shortest path from start to end with Bellman-Ford algorithm (Suurballe`s algorithm).
// found path state will be saved on Parents (needed to check from End Room)
// returns true if new path found, otherwise false
func searchShortPath(terrain *anthill) bool {
	usableRoomsQueue := &sortedQueue{}
	startRoom := terrain.Rooms[terrain.Start]
	endRoom := terrain.Rooms[terrain.End]
	startRoom.VisitIn, startRoom.VisitOut = true, true
	usableRoomsQueue.Enqueue(startRoom, 0, true)
	for usableRoomsQueue.Front != nil && !(endRoom.VisitIn || endRoom.VisitOut) {
		current := usableRoomsQueue.Dequeue()
		currentRoom := current.Room
		for next, value := range currentRoom.Paths {
			if value == BLOCKED || (!current.Mark && value == STABLE) {
				continue
			}
			addNext(currentRoom, next, current.Weight, value, usableRoomsQueue)
		}
	}
	isFind := endRoom.VisitIn || endRoom.VisitOut
	if isFind {
		replaceEdges(startRoom, endRoom)
		// clear flags
		for _, value := range terrain.Rooms {
			if value.VisitIn || value.VisitOut {
				value.ParentIn, value.ParentOut = nil, nil
				value.VisitIn, value.VisitOut = false, false
				value.Weight[0], value.Weight[1] = 0, 0
			}
		}
		startRoom.Separated = false
		endRoom.Separated = false
		usableRoomsQueue = nil
	}
	return isFind
}

// addNext - add into usableRoomsQueue next room with following rules
func addNext(cur, next *room, weight, state int, usableRoomsQueue *sortedQueue) {
	// if next room isn't visited then add without checking weights
	if !(next.VisitIn || next.VisitOut) {
		// we'll check next room for using on previous paths (separated flag)
		if next.Separated {
			next.VisitIn = true
			next.ParentIn = cur
			next.Weight[1] = weight + state
			// if it's usually path between nodes then add only in_node (check Surrballe's algo)
			if state == STABLE {
				usableRoomsQueue.Enqueue(next, next.Weight[1], false)
				return
			}
		}
		next.VisitOut = true
		next.ParentOut = cur
		next.Weight[0] = weight + state
		usableRoomsQueue.Enqueue(next, next.Weight[0], true)
		return
	}
	if !next.Separated {
		if weight+state >= next.Weight[0] {
			return
		}
		next.ParentOut = cur
		next.Weight[0] = weight + state
		usableRoomsQueue.Enqueue(next, next.Weight[0], true)
		return
	}
	if state == STABLE {
		if next.VisitIn && weight+state >= next.Weight[1] {
			return
		}
		next.VisitIn = true
		next.ParentIn = cur
		next.Weight[1] = weight + state
		usableRoomsQueue.Enqueue(next, next.Weight[1], false)
		return
	}
	if (next.VisitIn && weight+state < next.Weight[1]) || !next.VisitIn {
		next.VisitIn = true
		next.ParentIn = cur
		next.Weight[1] = weight + state
		usableRoomsQueue.Enqueue(next, next.Weight[1], false)
	}
	if (next.VisitOut && weight+state < next.Weight[0]) || !next.VisitOut {
		next.VisitOut = true
		next.ParentOut = cur
		next.Weight[0] = weight + state
		usableRoomsQueue.Enqueue(next, next.Weight[0], true)
	}
}

// replaceEdges - replace edges for finded paths. (Suurballe`s algorithm)
func replaceEdges(startRoom, endRoom *room) {
	r := endRoom
	for r != startRoom {
		var parent *room
		if r.ParentOut != nil && r.ParentIn != nil {
			i := 0
			for _, value := range r.Paths {
				if value == BLOCKED {
					i++
				}
			}
			if i > 1 {
				if r.Paths[r.ParentOut] == BLOCKED {
					parent = r.ParentOut
				} else {
					parent = r.ParentIn
				}
			} else {
				if r.Paths[r.ParentIn] == STABLE {
					parent = r.ParentIn
				} else {
					parent = r.ParentOut
				}
			}
		} else if r.ParentOut != nil {
			parent = r.ParentOut
		} else {
			parent = r.ParentIn
		}
		// reversing
		if r.Paths[parent] == STABLE {
			parent.Separated = true
			r.Separated = true
			r.Paths[parent] = REVERSED
			parent.Paths[r] = BLOCKED
		} else {
			parent.Separated = false
			r.Paths[parent] = STABLE
			parent.Paths[r] = STABLE
		}
		r = parent
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
	curStepsCount, used := fastCalcSteps(terrain.AntsCount, newPaths)
	// For debug
	// fmt.Printf("Steps: %d\n", curStepsCount)
	// for i := range newPaths {
	// 	start := newPaths[i].Front
	// 	fmt.Printf("Len: %d | %s", newPaths[i].Len, start.Room.Name)
	// 	for start.Room != endRoom {
	// 		start = start.Next
	// 		fmt.Printf(" --> %s", start.Room.Name)
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()
	if terrain.StepsCount == 0 || (terrain.StepsCount >= curStepsCount && used) {
		terrain.StepsCount = curStepsCount
		terrain.Result.Paths = newPaths
		return curStepsCount != 1
	}
	return false
}

// fastCalcSteps - calculate steps for paths and ants count
func fastCalcSteps(ants int, paths []*list) (int, bool) {
	steps, lossPerStep := 0, 0
	max, maxUsed := 0, false
	comingAnts := make(map[int]int)
	for _, value := range paths {
		comingAnts[value.Len]++
		if max < value.Len {
			max = value.Len
		}
	}
	if comingAnts[1] > 0 {
		return 1, true
	}
	for ants > 0 {
		steps++
		ants -= lossPerStep
		if steps == max && ants >= comingAnts[max] {
			maxUsed = true
		}
		lossPerStep += comingAnts[steps]
		ants -= comingAnts[steps]
	}
	return steps, maxUsed
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
		}
	}
	return steps, result
}
