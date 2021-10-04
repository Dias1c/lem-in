package anthill

import (
	"errors"
	"fmt"
)

type anthill struct {
	AntsCount  int              // Count Ants in anthill
	StepsCount int              // Count Steps for send all Ants to End Room
	Start, End string           // Start, End Room Names
	Rooms      map[string]*room // map["RoomName"]*Room
	Paths      []*path          // Paths
}

// getTerrainFromLines = Constructor.
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
	return &anthill{AntsCount: countAnts, Rooms: rooms, Start: startRoom, End: endRoom, Paths: []*path{}}, nil
}

// Validate - Checks anthill for correct. (I know just want to check all in 1 func)
func (a *anthill) Validate() error {
	// Check main params
	if a == nil {
		return fmt.Errorf("anthill can't be Null")
	} else if a.AntsCount <= 0 {
		return fmt.Errorf("incorrect count of ants on anthill: %q but expected >= 1", a.AntsCount)
	} else if a.Rooms == nil {
		return fmt.Errorf("rooms can't be Null")
	} else if len(a.Rooms) < 2 {
		return fmt.Errorf("in anthill should be minimum 2 rooms")
	} else if a.Start == "" {
		return fmt.Errorf("start room name can't be empty")
	} else if a.End == "" {
		return fmt.Errorf("end room name can't be empty")
	} else if _, isExist := a.Rooms[""]; isExist {
		return fmt.Errorf("room names can't be ''")
	}
	// Check Rooms For positions
	sliceRooms := make([]*room, len(a.Rooms))
	i := 0
	// Set map to slice for easy check + check paths, costs for nil
	for _, room := range a.Rooms {
		if room == nil {
			return fmt.Errorf("room can't be Null")
		} else if room.Name == "" {
			return fmt.Errorf("room name can't be empty")
		} else if room.Paths == nil {
			return fmt.Errorf("room paths can't be Null (RoomName: %v)", room.Name)
		} else if len(room.Paths) == 0 {
			return fmt.Errorf("there is no way to any room in the room (RoomName: %v)", room.Name)
		} else if _, isExist := room.Paths[room.Name]; isExist {
			return fmt.Errorf("room cannot have paths leading to itself (RoomName: %v)", room.Name)
		} else if room.Costs == nil {
			return fmt.Errorf("room path costs can't be Null (RoomName: %v)", room.Name)
		} else if len(room.Costs) != len(room.Paths) {
			return fmt.Errorf("room paths and costs sizes not equal (RoomName: %v)", room.Name)
		}
		sliceRooms[i] = room
		i++
	}
	// Check Room paths + positions
	for i, room := range sliceRooms {
		for _, innerRoom := range sliceRooms[i+1:] {
			if room.X == innerRoom.X && room.Y == innerRoom.Y {
				return fmt.Errorf("the rooms should not be on the same coordinates (1st room: %+v, 2nd room: %+v)", *room, *innerRoom)
			}
			// Check Room Paths
			if r, isExist := room.Paths[innerRoom.Name]; isExist {
				if r == nil {
					return fmt.Errorf("room can't be Null")
				}
				r, isExistTo := innerRoom.Paths[room.Name]
				if !isExistTo {
					return fmt.Errorf("the room you entered must have a way back")
				} else if r == nil {
					return fmt.Errorf("room can't be Null")
				}
			} else if r, isExist := innerRoom.Paths[room.Name]; isExist {
				if r == nil {
					return fmt.Errorf("room can't be Null")
				}
				r, isExistTo := room.Paths[innerRoom.Name]
				if !isExistTo {
					return fmt.Errorf("the room you entered must have a way back")
				} else if r == nil {
					return fmt.Errorf("room can't be Null")
				}
			}
		}
	}
	return nil
}

// Match - To Do
func (a *anthill) Match() (string, error) {
	if !a.HasItPath() {
		return "", errors.New("paths not found")
	}
	if err := a.SetBestPathsForCountAnts(); err != nil {
		return "", err
	}
	return fmt.Sprintf("Paths:\n%+v\n", a.Paths), nil
}

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

func (a *anthill) SetBestPathsForCountAnts() error {
	usableRooms, err := a.GetUsableRooms()
	if err != nil {
		return err
	}
	fmt.Println("SetBestPathsForCountAnts")
	stepsCnt := 0
	// Start Match and Find Count Steps + Best Paths
	for path := FindOneShortestPathByCost(a.Rooms[a.Start], a.Rooms[a.End], usableRooms); path != nil; path = FindOneShortestPathByCost(a.Rooms[a.Start], a.Rooms[a.End], usableRooms) {
		a.Paths = append(a.Paths, path)
		err := organizeIndependentPaths(a.Paths)
		if err != nil {
			return err
		}
		stepsCnt = getCountStepsForAllPaths(a.Paths, a.AntsCount)
		fmt.Printf("StepsCount: %d, for %d paths\n", stepsCnt, len(a.Paths))
		// if a.AntsCount != 0 && stepsCnt >= a.AntsCount {
		// 	break
		// }
	}

	return nil
}
