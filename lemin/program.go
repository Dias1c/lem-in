package lemin

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func getFileContent(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// getCountAntsFromLine - returns count ants, Changes size if line
func getCountAntsFromLine(lines *[]string) (int, error) {
	err := errors.New("getCountAntsFromLine: Lines == nil")
	if lines == nil {
		return 0, err
	} else if len(*lines) < 1 {
		return 0, errors.New("invalid number of Ants")
	}
	startIdx := 0
	countAnts := -1
	for i, line := range *lines {
		startIdx = i + 1
		if strings.HasPrefix(line, "#") {
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
// Checkking room names for doubles.
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
	// Gets All Rooms
	for i := 0; i < size; i++ {
		line := (*lines)[i]
		startIdx = i + 1
		if strings.HasPrefix(line, "#") { // Check for comment or Start|End Rooms
			isStart := false
			if line == "##start" {
				isStart = true
			}
			switch line {
			case "##start", "##end": // check for start or end room
				i++
				startIdx++
				if i < size {
					line = (*lines)[i]
					room := GetRoomFromLine(line)
					if room == nil {
						return nil, "", "", errors.New("invalid rooms for start or end not found")
					} else if isStart && startRoom == "" {
						startRoom = room.Name
					} else if !isStart && endRoom == "" {
						endRoom = room.Name
					} else {
						return nil, "", "", errors.New("there can be only 1 starting and 1 ending rooms on the terrain")
					}
				} else {
					return nil, "", "", errors.New("invalid rooms for start or end not found")
				}
			default:
				continue
			}
		}
		// Take rooms or break if is not valid room
		room := GetRoomFromLine(line)
		if room == nil {
			startIdx = startIdx - 1
			break
		}
		if _, isRoomExist := rooms[room.Name]; isRoomExist {
			return nil, "", "", errors.New("the names of the rooms should not be repeated")
		}
		rooms[room.Name] = room
	}
	if startRoom == "" || endRoom == "" {
		return nil, "", "", errors.New("invalid rooms for start or end not found")
	}
	*lines = (*lines)[startIdx:]
	return rooms, startRoom, endRoom, nil
}

// setPathsFromLines - set
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
	for i, line := range *lines {
		startIdx = i + 1
		if strings.HasPrefix(line, "#") {
			continue
		} else {
			if isCorrect := pattern.MatchString(line); isCorrect {
				submatch := pattern.FindStringSubmatch(line)
				//Set Paths
				nameFrom := submatch[1]
				nameTo := submatch[2]
				fmt.Printf("%v - %v\n", nameFrom, nameTo)
				if _, roomFound := rooms[nameFrom]; !roomFound {
					return fmt.Errorf("unknown room name: %q", nameFrom)
				} else if _, roomFound := rooms[nameTo]; !roomFound {
					return fmt.Errorf("unknown room name: %q", nameTo)
				}

				// Реализовать добавление путей для комнат + проверка
			} else {
				return fmt.Errorf("invalid path: %q", line)
			}
		}
	}
	*lines = (*lines)[startIdx:]
	return nil
}
