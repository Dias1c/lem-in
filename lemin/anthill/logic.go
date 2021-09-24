package anthill

import (
	"container/list"
	"fmt"
)

func (a *anthill) HasItPath() bool {
	moved := map[string]bool{a.Start: true}
	stack := []*room{a.Rooms[a.Start]}
	stackSize := 1

	for stackSize > 0 {
		for name := range stack[0].Paths {
			if name == a.End {
				return true
			} else if !moved[name] {
				stack = append(stack, a.Rooms[name])
				stackSize++
			}
		}
		stack = stack[1:]
		stackSize--
	}
	return false
}

func (a *anthill) GetUsableRooms() (map[string]bool, error) {
	//Get all using Rooms from start to end and from end to start
	startEnd, err := a.GetUsingRoomNamesFromTo(a.Start, a.End)
	if err != nil {
		return nil, err
	}
	endStart, err := a.GetUsingRoomNamesFromTo(a.End, a.Start)
	if err != nil {
		return nil, err
	}
	// Set Rooms wich in Start + End
	result := make(map[string]bool, len(startEnd))
	removedRooms := make(map[string]bool) // Delete it
	for name := range startEnd {
		if endStart[name] {
			result[name] = true
		} else {
			removedRooms[name] = true
		}
	}
	result[a.Start] = true
	result[a.End] = true
	fmt.Printf("Removed Roooms: %+v\n", removedRooms)
	return result, nil
}

func (a *anthill) GetUsingRoomNamesFromTo(from, to string) (map[string]bool, error) {
	if a.Rooms[from] == nil || a.Rooms[to] == nil {
		return nil, fmt.Errorf("GetUsingRoomNamesFromTo: RoomFrom: %s (%+v) or RoomTo: %s (%+v) not found!", from, a.Rooms[from], to, a.Rooms[to])
	}
	countRooms := len(a.Rooms)
	moved := make(map[string]bool, countRooms)
	moved[from] = true
	moved[to] = true
	// Start Match
	stack := []*room{a.Rooms[from]}
	remains := 1
	for remains > 0 {
		for name := range stack[0].Paths {
			if !moved[name] {
				stack = append(stack, a.Rooms[name])
				remains++
				moved[name] = true
			}
		}
		stack = stack[1:]
		remains--
	}
	delete(moved, from)
	delete(moved, to)
	return moved, nil
}

func FindOneShortestPathByCost(from, to *room, usableRooms map[string]bool) *list.List {
	stack, stackSize := []*room{from}, 1
	movedRooms := make(map[string]bool, 2)
	movedRooms[from.Name] = true

	isFind := false
	for stackSize > 0 && !isFind {
		curRoom := stack[0]
		for name, room := range curRoom.Paths {
			if usableRooms[name] && !movedRooms[name] && curRoom.Costs[name] == 1 {
				room.PrevRoom = curRoom
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
	result := list.New()
	curRoom := to
	prevRoom := curRoom.PrevRoom
	for curRoom != from {
		result.PushFront(curRoom)
		prevRoom.Costs[curRoom.Name] = 0
		curRoom = prevRoom
		prevRoom = curRoom.PrevRoom
	}
	result.PushFront(curRoom)
	PrintRoomsInLinkedList(result)
	return result
}

func (a *anthill) SetBestPathsForCountAnts() error {
	usableRooms, err := a.GetUsableRooms()
	if err != nil {
		return err
	}
	fmt.Println("SetBestPathsForCountAnts")
	// Start Match
	FindOneShortestPathByCost(a.Rooms[a.Start], a.Rooms[a.End], usableRooms)
	FindOneShortestPathByCost(a.Rooms[a.Start], a.Rooms[a.End], usableRooms)
	FindOneShortestPathByCost(a.Rooms[a.Start], a.Rooms[a.End], usableRooms)
	// Set Paths
	// Match CountSteps
	return nil
}
