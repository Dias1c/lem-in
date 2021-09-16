package lemin

import (
	"errors"
	"fmt"
)

type terrain struct {
	AntsCount  int              // Count Ants in Terrain
	Start, End string           // Start, End Room Names
	Rooms      map[string]*room // map["RoomName"]*Room
	Paths      []*paths         // Paths
}

// getTerrainFromLines = Constructor.
func getTerrainFromLines(lines []string) (*terrain, error) {
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
	return &terrain{AntsCount: countAnts, Rooms: rooms, Start: startRoom, End: endRoom, Paths: []*paths{}}, nil
}

// IsCorrect - Checks terrain for correct.
func (t *terrain) Validate() error {
	// Check main params
	if t == nil {
		return fmt.Errorf("terrain can't be Null")
	} else if t.AntsCount <= 0 {
		return fmt.Errorf("incorrect count of ants on terrain: %q but expected >= 1", t.AntsCount)
	} else if t.Rooms == nil {
		return fmt.Errorf("rooms can't be Null")
	} else if len(t.Rooms) < 2 {
		return fmt.Errorf("in terrain should be minimum 2 rooms")
	} else if t.Start == "" {
		return fmt.Errorf("start room name can't be empty")
	} else if t.End == "" {
		return fmt.Errorf("end room name can't be empty")
	} else if _, isExist := t.Rooms[""]; isExist {
		return fmt.Errorf("room names can't be ''")
	}
	// Check Rooms For positions
	sliceRooms := make([]*room, len(t.Rooms))
	i := 0
	// Set map to slice for easy check + check paths for nil
	for _, room := range t.Rooms {
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
func (t *terrain) Match() (string, error) {
	if t.Paths == nil {
		t.Paths = []*paths{}
	}
	levels := t.GetLevels()
	PrintLevels(levels)
	t.RemoveUselessFromLevels(levels)
	PrintLevels(levels)
	if len(t.Paths) == 0 {
		return "", errors.New("Paths NOT FOUND")
	}
	return fmt.Sprintf("Paths:\n%+v\n", t.Paths), nil
}
