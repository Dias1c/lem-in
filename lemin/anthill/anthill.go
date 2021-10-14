package anthill

import (
	"errors"
	"strings"
)

// CreateAnthill - returns anthill by default data
func CreateAnthill() *anthill {
	result := &anthill{}
	result.Rooms = make(map[string]*room)
	result.FieldInfo = &fieldInfo{UsingCoordinates: make(map[int]map[int]bool)}
	result.Result = &Result{}
	return result
}

// Set nil data for anthill
func (a *anthill) ClearData() {
	a.Rooms = nil
	a.FieldInfo = nil
}

//
func (a *anthill) ReadDataFromLine(line string) error {
	if line == "" || strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
		return nil
	}
	switch a.FieldInfo.MODE {
	case FIELD_PATHS:
		err := a.SetPathsFromLine(line)
		if err != nil {
			return err
		}
	case FIELD_ROOMS:
		if strings.HasPrefix(line, "##") {
			if line == "##start" && !a.FieldInfo.Start && !a.FieldInfo.IsStart && !a.FieldInfo.IsEnd {
				a.FieldInfo.IsStart = true
				return nil
			} else if line == "##end" && !a.FieldInfo.End && !a.FieldInfo.IsEnd && !a.FieldInfo.IsStart {
				a.FieldInfo.IsEnd = true
				return nil
			}
			return errors.New("error with ## command")
		}
		if a.FieldInfo.IsStart || a.FieldInfo.IsEnd {
			err := a.SetMainRooms(line, a.FieldInfo.IsStart)
			if err != nil {
				return err
			}
			if a.FieldInfo.IsStart {
				a.FieldInfo.IsStart = false
				a.FieldInfo.Start = true
			} else {
				a.FieldInfo.IsEnd = false
				a.FieldInfo.End = true
			}
			return err
		} else if len(strings.Split(line, " ")) != 3 {
			a.FieldInfo.MODE = FIELD_PATHS
			a.FieldInfo.UsingCoordinates = nil
		} else {
			_, err := a.SetRoomFromLine(line)
			return err
		}
	case FIELD_ANTS:
		err := a.SetAntsFromLine(line)
		if err != nil {
			return err
		}
		a.FieldInfo.MODE = FIELD_ROOMS
	}
	return nil
}

// Match -
func (a *anthill) Match() error {
	for {
		if !searchShortPath(a) {
			// path not found, then check for prev path count
			if a.StepsCount > 0 {
				return nil
			} else {
				return errors.New("path not found")
			}
		}
		// checking effective of new short path
		// current steps count < previous steps count
		// if effective then replace result to new (returns true)
		// if not then return previous result (returns false)
		if !checkEffective(a) {
			return nil
		}
	}
}
