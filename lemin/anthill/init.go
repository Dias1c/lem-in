package anthill

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Rules for CountAnts:
// CountAnts must be > 0

// SetAntsFromLine - set data about count ants from line to anthill
func (a *anthill) SetAntsFromLine(line string) error {
	countAnts, err := strconv.Atoi(line)
	if err != nil || countAnts < 1 {
		return errors.New("invalid number of Ants")
	}
	a.AntsCount = countAnts
	a.Result.AntsCount = countAnts
	return nil
}

// Rules for Room:
// Names must be unique
// Coordinates must be unique

// SetRoomFromLine - insert rooms into anthill, returns error if invalid room
func (a *anthill) SetRoomFromLine(line string) (*room, error) {
	splited := strings.Split(line, " ")
	if len(splited) != 3 || len(splited[0]) < 1 {
		return nil, errors.New("invalid format of room")
	} else if strings.HasPrefix(splited[0], "L") {
		return nil, errors.New("room name can't be started with 'L'")
	} else if strings.Contains(splited[0], "-") {
		return nil, errors.New("room name can't have '-'")
	} else if _, ok := a.Rooms[splited[0]]; ok {
		return nil, fmt.Errorf("room name duplicated: '%v'", splited[0])
	}

	name := splited[0]
	x, errX := strconv.Atoi(splited[1])
	y, errY := strconv.Atoi(splited[2])
	if errX != nil || errY != nil {
		return nil, errors.New("room coords can only be numbers")
	} else if _, ok := a.FieldInfo.UsingCoordinates[x]; ok {
		if a.FieldInfo.UsingCoordinates[x][y] {
			return nil, fmt.Errorf("room coords must be unique; room name: '%v'", name)
		}
	} else {
		a.FieldInfo.UsingCoordinates[x] = make(map[int]bool)
	}
	a.FieldInfo.UsingCoordinates[x][y] = true

	room := &room{
		Name:   name,
		X:      x,
		Y:      y,
		Paths:  make(map[*room]int),
		Weight: [2]int{0, 0},
	}
	a.Rooms[name] = room
	return room, nil
}

// SetMainRooms - insert rooms into anthill and set Start or End by marker startOrEnd
func (a *anthill) SetMainRooms(line string, startOrEnd bool) error {
	room, err := a.SetRoomFromLine(line)
	if err != nil {
		return err
	}
	if startOrEnd {
		a.Start = room.Name
	} else {
		a.End = room.Name
	}
	return nil
}

// Rules for Room Relations
// Room cant has path to themseld

// SetPathsFromLine - builds relationships between rooms available in anthill by line;
func (a *anthill) SetPathsFromLine(line string) error {
	splited := strings.Split(line, "-")
	if len(splited) != 2 || len(splited[0]) < 1 || len(splited[1]) < 1 {
		return errors.New("invalid format of path")
	}
	name1, name2 := splited[0], splited[1]
	if name1 == name2 {
		return fmt.Errorf("rooms can't link themselves. Line: '%v'", line)
	}
	room1 := a.Rooms[name1]
	room2 := a.Rooms[name2]
	if room1 == nil || room2 == nil {
		return fmt.Errorf("path contains unknown room. Line: '%v'", line)
	}
	// if _, ok := room1.Paths[room2]; ok {
	// 	return fmt.Errorf("rooms already linked. Line: '%v'", line)
	// }
	room1.Paths[room2] = STABLE
	room2.Paths[room1] = STABLE
	return nil
}
