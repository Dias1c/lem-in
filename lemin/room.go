package lemin

import (
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
}

// GetRoomFromLine - Returns nil if line incorrect
func GetRoomFromLine(line string) *room {
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
	}
}
