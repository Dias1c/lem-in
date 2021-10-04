package anthill

import (
	"fmt"
)

func FindOneShortestPathByCost(from, to *room, usableRooms map[string]bool) *path {
	stack, stackSize := []*room{from}, 1
	prevRooms := make(map[*room]*room, len(usableRooms))
	movedRooms := make(map[string]bool, len(usableRooms))
	movedRooms[from.Name] = true

	isFind := false
	for stackSize > 0 && !isFind {
		curRoom := stack[0]
		var wantedCost int8 = 1
		if curRoom.PrevRoom != nil {
			wantedCost = -1
			if curRoom.PrevRoom != nil {
				curRoom.PrevRoom.PrevRoom = nil
			}
		}
		for name, room := range curRoom.Paths {
			if usableRooms[name] && !movedRooms[name] && curRoom.Costs[name] == wantedCost {
				// room.PrevRoom = curRoom
				prevRooms[room] = curRoom
				stack = append(stack, room)
				stackSize++
				if room == to {
					isFind = true
					break
				}
				movedRooms[name] = true
			}
		}
		stack = stack[1:]
		stackSize--
	}
	if !isFind {
		return nil
	}
	result := &path{}
	result.PushFront(to)
	curRoom := prevRooms[to]
	// prevRoom := prevRooms[curRoom]
	for curRoom != nil {
		curRoom.Costs[result.Front.Room.Name] = 0
		result.Front.Room.Costs[curRoom.Name] = -1
		result.Front.Room.PrintRoomCosts()
		curRoom.PrevRoom = prevRooms[curRoom]
		result.PushFront(curRoom)

		// curRoom.PrintRoomCosts()
		curRoom = curRoom.PrevRoom
	}
	result.Front.Room.PrintRoomCosts()
	result.RemoveFront()

	fmt.Println("Finded new Path")
	PrintRoomsInLinkedList(result)
	return result
}

func getCountStepsForAllPaths(paths []*path, antsCount int) int {
	steps := 0
	cntAntsOnFinish := make(map[int]int, len(paths))
	for _, path := range paths {
		cntAntsOnFinish[path.Len-1] += 1
	}
	if _, isHave := cntAntsOnFinish[0]; isHave {
		return 1
	}
	curFinishFlow := 0
	for ; antsCount > 0; steps++ {
		curFinishFlow += cntAntsOnFinish[steps]
		antsCount -= curFinishFlow
	}
	return steps
}

func organizeIndependentPaths(paths []*path) error {
	if len(paths) < 2 {
		return nil
	}
	fmt.Println("organizeIndependentPaths: Before")
	for _, v := range paths {
		PrintRoomsInLinkedList(v)
	}
	for i := 0; i < len(paths); i++ {
		if paths[i].Len < 2 {
			return fmt.Errorf("organizeIndependentPaths: paths[%d].len() < 2", i)
		}
		iPrev := paths[i].Front
		iCur := paths[i].Front.Next
		for iCur != nil && iCur.Next != nil {
			ipRoom := iPrev.Room
			icRoom := iCur.Room

			if icRoom.Costs[ipRoom.Name] == 0 {
				fmt.Printf("I := %v\niPrevRoom: %+v\niCurRoom: %+v\n", i, ipRoom, icRoom)
				isSwapped := false
				for j := 0; j < len(paths) && !isSwapped; j++ {
					if j == i {
						continue
					}
					jPrev := paths[j].Front
					jCur := paths[j].Front.Next
					for jCur != nil && jCur.Next != nil {
						if iPrev.Room == jCur.Room && iCur.Room == jPrev.Room {
							iPrev.Next, jPrev.Next = jCur.Next, iCur.Next

							paths[i].RefreshData()
							paths[j].RefreshData()
							PrintRoomsInLinkedList(paths[i])
							PrintRoomsInLinkedList(paths[j])
							isSwapped = true
							fmt.Printf("Swapped: %v\n", isSwapped)
							break
						}
						jPrev = jCur
						jCur = jCur.Next
					}
				}
				fmt.Printf("Swapped: %v\n", isSwapped)
				if !isSwapped {
					return fmt.Errorf("organizeIndependentPaths: for paths[%d] data not sapped")
				}
				icRoom.Costs[ipRoom.Name] = 1
				fmt.Printf("Break\n")
				break
			}
			iPrev = iCur
			iCur = iCur.Next
		}
	}

	fmt.Println("organizeIndependentPaths: After")
	for _, v := range paths {
		PrintRoomsInLinkedList(v)
	}
	return nil
}

func getPathsCopyWithOrigRooms(paths []*path) []*path {
	result := make([]*path, len(paths))
	for i := 0; i < len(result); i++ {
		result[i] = paths[i].GetCopyWithOrigRoom()
	}
	return result
}
