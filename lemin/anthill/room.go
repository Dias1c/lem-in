package anthill

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const (
	PatternRoomName = `[.\dA-KM-Za-zА-Яа-я]{1}[.\d\wА-Яа-я]{0,}`
)

//
type room struct {
	Name      string           // RoomName
	AntsCount int              // Count Ants in room
	X, Y      int              // Coordinates
	Paths     map[string]*room // Adjacent rooms
	Costs     map[string]byte  // Cost For Every Room in Path
	PrevRoom  *room            // From Wich room did you come?
}

// GetRoomFromLine - Returns nil if line incorrect
func getRoomFromLine(line string) *room {
	pattern, err := regexp.Compile(fmt.Sprintf(`^(%v) (\d{1,}) (\d{1,})$`, PatternRoomName))
	if err != nil {
		log.Printf("GetRoomFromLine: %v", err)
		return nil
	} else if isMatch := pattern.MatchString(line); !isMatch {
		return nil
	}
	submatch := pattern.FindStringSubmatch(line)
	name := submatch[1]
	coorx, errX := strconv.Atoi(submatch[2])
	coory, errY := strconv.Atoi(submatch[3])
	if errX != nil || errY != nil {
		return nil
	}
	return &room{
		Name:  name,
		X:     coorx,
		Y:     coory,
		Paths: make(map[string]*room, 1),
		Costs: make(map[string]byte, 1),
	}
}

// addPath - adds paths for both rooms
func addPath(rooms map[string]*room, from, to string) error {
	if rooms == nil {
		return errors.New("addPath: rooms is nil")
	}
	roomFrom, isFromExist := rooms[from]
	roomTo, isToExist := rooms[to]
	if !isFromExist || !isToExist {
		return fmt.Errorf("invalid room names on paths: %q-%q", from, to)
	}
	rooms[from].Paths[to] = roomTo
	rooms[to].Paths[from] = roomFrom

	rooms[from].Costs[to] = 1
	rooms[to].Costs[from] = 1
	return nil
}
