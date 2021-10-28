package anthill

import (
	"log"
)

// SearchShortPath - search shortest path from start to end with Bellman-Ford algorithm (Suurballe`s algorithm).
// found path state will be saved on UsingOnPath, Parent (needed to check from End Room)
// returns true if found, otherwise false
func searchShortPath(terrain *anthill) bool {
	usableRoomsQueue := &sortedQueue{}
	startRoom := terrain.Rooms[terrain.Start]
	endRoom := terrain.Rooms[terrain.End]
	startRoom.VisitIn, startRoom.VisitOut = true, true
	usableRoomsQueue.Enqueue(startRoom, 0, true)
	// i := 0
	for usableRoomsQueue.Front != nil && !(endRoom.VisitIn || endRoom.VisitOut) {
		current := usableRoomsQueue.Dequeue()
		currentRoom := current.Room
		// fmt.Printf("Current: %s\n", current.Name)
		for next, value := range currentRoom.Paths {
			if value == BLOCKED || (!current.Mark && value == STABLE) {
				continue
			}
			// str := "BLOCKED"
			// if value == STABLE {
			// 	str = "STABLE"
			// } else if value == REVERSED {
			// 	str = "REVERSED"
			// }
			// fmt.Printf("Adding %s from %s with %s\n", next.Name, current.Name, str)
			addNext(currentRoom, next, current.Weight, value, usableRoomsQueue)
		}
		// if i == 1000 {
		// 	fmt.Print("Queue state: ")
		// 	usableRoomsQueue.DebugPrint()
		// 	fmt.Println()
		// 	os.Exit(3)
		// }
		// i++
		// fmt.Println(i)
		// fmt.Print("Queue state: ")
		// usableRoomsQueue.DebugPrint()
		// fmt.Println()
	}
	isFind := endRoom.VisitIn || endRoom.VisitOut
	// fmt.Print("Queue state: ")
	// usableRoomsQueue.DebugPrint()
	// fmt.Println()
	// debugVisitedRooms(visitedRooms)
	if isFind {
		// fmt.Println("### Debug Visited Rooms ###")
		// for _, value := range terrain.Rooms {
		// 	if value.VisitIn || value.VisitOut {
		// 		parentInName, parentOutName := "nil", "nil"
		// 		if value.ParentIn != nil {
		// 			parentInName = value.ParentIn.Name
		// 		}
		// 		if value.ParentOut != nil {
		// 			parentOutName = value.ParentOut.Name
		// 		}
		// 		fmt.Printf("Room: %s\n\tParent: o-%s, i-%s\n\tWeight: o-%d, i-%d\n", value.Name, parentOutName, parentInName, value.Weight[0], value.Weight[1])
		// 	}
		// }
		// fmt.Println("###########################\n")
		// to do replace edges
		replaceEdges(startRoom, endRoom)
		// clear flags
		for _, value := range terrain.Rooms {
			if value.VisitIn || value.VisitOut {
				value.ParentIn, value.ParentOut = nil, nil
				value.VisitIn, value.VisitOut = false, false
				value.Weight[0], value.Weight[1] = 0, 0
				value.InNewPath = false
			}
		}
	}
	startRoom.Separated = false
	endRoom.Separated = false
	return isFind
}

// func debugVisitedRooms(visitedRooms map[*room]bool) {
// 	fmt.Println("### Debug Visited Rooms ###")
// 	i := 1
// 	for key := range visitedRooms {
// 		parentName := "nil"
// 		if key.Parent != nil {
// 			parentName = key.Parent.Name
// 		}
// 		fmt.Printf("%d. Room: %s\n\tParent: %s\n\tWeight: %d\n", i, key.Name, parentName, key.Weight)
// 		i++
// 	}
// 	fmt.Println("###########################\n")
// }

// addNext - add into usableRoomsQueue next room
func addNext(cur, next *room, weight, state int, usableRoomsQueue *sortedQueue) {
	if !(next.VisitIn || next.VisitOut) {
		if next.Separated {
			next.VisitIn = true
			next.ParentIn = cur
			next.Weight[1] = weight + state
			if state == STABLE {
				usableRoomsQueue.Enqueue(next, next.Weight[1], false)
				return
			}
		}
		next.VisitOut = true
		next.ParentOut = cur
		next.Weight[0] = weight + state
		usableRoomsQueue.Enqueue(next, next.Weight[0], true)
		// fmt.Printf("New room added: %s|%d\n", next.Name, weight)
		return
	}
	if !next.Separated {
		if weight+state >= next.Weight[0] {
			return
		}
		next.ParentOut = cur
		next.Weight[0] = weight + state
		usableRoomsQueue.Enqueue(next, next.Weight[0], true)
	} else {
		if state == STABLE {
			if next.VisitIn && weight+state >= next.Weight[1] {
				return
			}
			next.VisitIn = true
			next.ParentIn = cur
			next.Weight[1] = weight + state
			usableRoomsQueue.Enqueue(next, next.Weight[1], false)
		} else {
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
	}
	// if next.Separated && !next.Marked && cur.Paths[next] == REVERSED {
	// 	if next.Weight < weight {
	// 		next.Marked = true
	// 	} else {
	// 		markNext(cur, next, weight, visitedRooms)
	// 	}
	// } else if next.Weight > weight {
	// 	markNext(cur, next, weight, visitedRooms)
	// } else {
	// 	if next.Separated {
	// 		fmt.Printf("Returned: %s|sep\n", next.Name)
	// 	} else {
	// 		fmt.Printf("Returned: %s\n", next.Name)
	// 	}
	// 	return
	// }
	// usableRoomsQueue.SortEnqueue(next)
	// fmt.Printf("Visited room added: %s|%d\n", next.Name, weight)
}

// replaceEdges - replace edges for finded paths. (Suurballe`s algorithm)
func replaceEdges(startRoom, endRoom *room) {
	r := endRoom
	// fmt.Printf("%s --> ", r.Name)
	for r != startRoom {
		var parent *room
		if r.ParentOut != nil && r.ParentIn != nil {
			i := 0
			for _, value := range r.Paths {
				if value == BLOCKED {
					i++
				}
			}
			if i != 1 {
				if r.Paths[r.ParentOut] == BLOCKED {
					parent = r.ParentOut
				} else {
					parent = r.ParentIn
				}
			} else {
				if r.Paths[r.ParentOut] == STABLE {
					parent = r.ParentOut
				} else {
					parent = r.ParentIn
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
		} else if r.Paths[parent] == REVERSED {
			log.Fatalf("path %s-%s is REVERSED", r.Name, parent.Name)
		} else {
			parent.Separated = false
			r.Paths[parent] = STABLE
			parent.Paths[r] = STABLE
		}
		r = parent
		// fmt.Printf("%s --> ", r.Name)
	}
	// fmt.Println()
	// i := 0
	// r := endRoom
	// fmt.Printf("%s --> ", r.Name)
	// for r != startRoom {
	// 	if r.Parent.Separated && r.Parent.Marked && r.Paths[r.Parent] == STABLE {
	// 		for key, value := range r.Parent.Paths {
	// 			if value == BLOCKED {
	// 				fmt.Printf("\nChange parent %s from %s to %s\n", r.Parent.Name, r.Parent.Parent.Name, key.Name)
	// 				r.Parent.Parent = key
	// 				break
	// 			}
	// 		}
	// 	}
	// 	if r.Parent.InNewPath {
	// 		setNewParent(r.Parent)
	// 	}
	// 	r.Parent.InNewPath = true
	// 	// reversing
	// 	if r.Paths[r.Parent] == STABLE {
	// 		r.Parent.Separated = true
	// 		r.Separated = true
	// 		r.Paths[r.Parent] = REVERSED
	// 		r.Parent.Paths[r] = BLOCKED
	// 	} else if r.Paths[r.Parent] == REVERSED {
	// 		log.Fatalf("path %s-%s is REVERSED", r.Name, r.Parent.Name)
	// 	} else {
	// 		r.Parent.Separated = false
	// 		r.Paths[r.Parent] = STABLE
	// 		r.Parent.Paths[r] = STABLE
	// 	}
	// 	r = r.Parent
	// 	fmt.Printf("%s --> ", r.Name)
	// 	if i == 100 {
	// 		os.Exit(2)
	// 	}
	// 	i++
	// }
	// fmt.Println()
}

// func setNewParent(current *room) {
// 	minWeight := MAX_INT // max int
// 	for key, value := range current.Paths {
// 		if value == STABLE {
// 			if key.Weight < minWeight {
// 				minWeight = key.Weight
// 				current.Parent = key
// 			}
// 		}
// 	}
// }

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
	// fmt.Println("### Debug #Paths ###")
	// for _, v := range newPaths {
	// 	fr := v.Front
	// 	for fr != v.Back {
	// 		fmt.Printf("%s -> ", fr.Room.Name)
	// 		fr = fr.Next
	// 	}
	// 	fmt.Printf("%s\n", fr.Room.Name)
	// }
	// fmt.Println("####################\n")
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
