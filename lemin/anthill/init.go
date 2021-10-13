package anthill

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// getCountAntsFromLine - returns count ants, Changes size if line
func getCountAntsFromLine(lines *[]string) (int, error) {
	err := errors.New("getCountAntsFromLine: Lines == nil")
	if lines == nil {
		return 0, err
	} else if len(*lines) < 1 {
		return 0, errors.New("invalid number of Ants")
	}
	startIdx := 0
	countAnts := 0
	for i, line := range *lines {
		startIdx = i + 1
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		} else {
			countAnts, err = strconv.Atoi(line)
			if err != nil || countAnts < 1 {
				return 0, errors.New("invalid number of Ants")
			} else {
				*lines = (*lines)[startIdx:]
				return countAnts, nil
			}
		}
	}
	*lines = (*lines)[startIdx:]
	return 0, errors.New("count Ants not found")
}

// getRoomsFromLine - returns Rooms, startRoom, endRoom, error if exists.
// Checkking room names and coordinates for doubles
func getRoomsFromLines(lines *[]string) (map[string]*room, string, string, error) {
	err := errors.New("getRoomsFromLine: Lines == nil")
	if lines == nil {
		return nil, "", "", err
	} else if len(*lines) < 2 {
		return nil, "", "", errors.New("not enough rooms")
	}
	startIdx := 0
	size := len(*lines)
	rooms, startRoom, endRoom := make(map[string]*room, 2), "", ""
	uniqueCoords := make(map[int]map[int]bool, len(*lines)/3)
	// Gets All Rooms
	for i := 0; i < size; i++ {
		line := (*lines)[i]
		startIdx = i + 1
		if strings.HasPrefix(line, "#") { // Check for comment or Start|End Rooms
			switch line {
			case "##start", "##end": // check for start or end room
				isStart := false
				if line == "##start" {
					isStart = true
				}
				i++
				startIdx++
				if i < size {
					line = (*lines)[i]
					room := getRoomFromLine(line)
					if room == nil {
						return nil, "", "", fmt.Errorf("invalid rooms for start or end. Problem with: '%v'", line)
					} else if isStart && startRoom == "" {
						startRoom = room.Name
					} else if !isStart && endRoom == "" {
						endRoom = room.Name
					} else {
						return nil, "", "", errors.New("there can be only 1 starting and 1 ending rooms on the anthill")
					}
				} else {
					return nil, "", "", errors.New("invalid rooms for start or end not found")
				}
			default:
				continue
			}
		} else if line == "" {
			continue
		}
		// Take rooms or break if is not valid room
		room := getRoomFromLine(line)
		if room == nil {
			startIdx = startIdx - 1
			break
		}
		// Check For Valid Room
		if _, isRoomExist := rooms[room.Name]; isRoomExist {
			return nil, "", "", fmt.Errorf("the names of the rooms should not be repeated. Room: '%v'", room.Name)
		}
		if uniqueCoords[room.X] != nil { // Check Unique Coordinates
			if uniqueCoords[room.X][room.Y] {
				return nil, "", "", fmt.Errorf("room coords must be unique. Room: '%v'", room.Name)
			}
		} else {
			uniqueCoords[room.X] = make(map[int]bool)
		}
		uniqueCoords[room.X][room.Y] = true

		rooms[room.Name] = room
	}
	if startRoom == "" || endRoom == "" {
		return nil, "", "", errors.New("invalid rooms for start or end not found")
	}
	*lines = (*lines)[startIdx:]
	return rooms, startRoom, endRoom, nil
}

// setPathsFromLines - set paths
func setPathsFromLines(lines *[]string, rooms map[string]*room) error {
	err := errors.New("SetPathsFromLines: Lines == nil")
	if lines == nil {
		return err
	} else if len(*lines) < 1 {
		return errors.New("setPathsFromLines: no connections")
	} else if rooms == nil {
		return errors.New("setPathsFromLines: no rooms")
	}
	startIdx := 0
	pattern, err := regexp.Compile(fmt.Sprintf(`^(%v)\-(%v)$`, PatternRoomName, PatternRoomName))
	if err != nil {
		return err
	}
	for i, line := range *lines {
		startIdx = i + 1
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		} else {
			if isCorrect := pattern.MatchString(line); isCorrect {
				submatch := pattern.FindStringSubmatch(line)
				// Set Paths
				nameFrom := submatch[1]
				nameTo := submatch[2]
				// fmt.Printf("%v - %v\n", nameFrom, nameTo)
				if _, roomFound := rooms[nameFrom]; !roomFound {
					return fmt.Errorf("unknown room name: %q", nameFrom)
				} else if _, roomFound := rooms[nameTo]; !roomFound {
					return fmt.Errorf("unknown room name: %q", nameTo)
				} else if nameFrom == nameTo {
					return fmt.Errorf("invalid path: %v-%v", nameFrom, nameTo)
				}

				addPath(rooms, nameFrom, nameTo)
			} else {
				return fmt.Errorf("invalid path: %q", line)
			}
		}
	}
	*lines = (*lines)[startIdx:]
	return nil
}
