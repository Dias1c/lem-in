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
		Name:     name,
		X:        coorx,
		Y:        coory,
		PathsIn:  make(map[string]*room, 1),
		PathsOut: make(map[string]*room, 1),
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
	rooms[from].PathsOut[to] = roomTo
	rooms[to].PathsOut[from] = roomFrom
	return nil
}
