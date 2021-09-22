package lemin

import (
	"fmt"
)

func (t *terrain) HasItPath() bool {
	moved := map[string]bool{t.Start: true}
	stack := []*room{t.Rooms[t.Start]}
	stackSize := 1

	for stackSize > 0 {
		for name := range stack[0].Paths {
			if name == t.End {
				return true
			} else if !moved[name] {
				stack = append(stack, t.Rooms[name])
				stackSize++
			}
		}
		stack = stack[1:]
		stackSize--
	}
	return false
}

func (t *terrain) GetLevels() []map[string]*room {
	moved := make(map[string]bool, len(t.Rooms))
	moved[t.Start] = true
	moved[t.End] = true
	stack := []*room{t.Rooms[t.Start]}
	levels := []map[string]*room{{t.Start: t.Rooms[t.Start]}}
	lvl := 0
	added := 1
	for len(stack) != 0 {
		levels = append(levels, make(map[string]*room))
		lvl++
		cutPos := 0
		for _, st := range stack[:added] {
			for name, room := range st.Paths {
				if !moved[name] {
					stack = append(stack, room)
					levels[lvl][name] = room
					moved[name] = true
					cutPos++
				}
			}
		}
		stack = stack[added:]
		added = cutPos
	}
	levels = levels[1:]
	// levels[lvl][t.End] = t.Rooms[t.End]
	if len(levels) >= 1 {
		levels = levels[:len(levels)-1]
	}
	return levels
}

func (t *terrain) GetUsingRoomNamesFromTo(from, to string) (map[string]bool, error) {
	if t.Rooms[from] == nil || t.Rooms[to] == nil {
		return nil, fmt.Errorf("GetUsingRoomNamesFromTo: RoomFrom: %s (%+v) or RoomTo: %s (%+v) not found!", from, t.Rooms[from], to, t.Rooms[to])
	}
	countRooms := len(t.Rooms)
	moved := make(map[string]bool, countRooms)
	moved[from] = true
	moved[to] = true
	// Start Match
	stack := []*room{t.Rooms[from]}
	remains := 1
	for remains > 0 {
		for name := range stack[0].Paths {
			if !moved[name] {
				stack = append(stack, t.Rooms[name])
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

func (t *terrain) RemoveUselessFromLevels(levels []map[string]*room) error {
	countLevels := len(levels)

	//Get all using Rooms from start to end and from end to start
	startEnd, err := t.GetUsingRoomNamesFromTo(t.Start, t.End)
	if err != nil {
		return err
	}
	endStart, err := t.GetUsingRoomNamesFromTo(t.End, t.Start)
	if err != nil {
		return err
	}

	removedRooms := make(map[string]bool)
	// Removes useless. Start removes from end
	for i := countLevels - 1; i > -1; i-- {
		level := levels[i]
		//Remove rooms from level
		for _, curRoom := range level {
			if len(curRoom.Paths) == 1 { // Remove if has only 1 path
				delete(level, curRoom.Name)
				fmt.Printf("Removed %q have only 1 path\n", curRoom.Name)
				removedRooms[curRoom.Name] = true
			} else if !startEnd[curRoom.Name] || !endStart[curRoom.Name] {
				delete(level, curRoom.Name)
				fmt.Printf("Removed %q Start To End or End to Start path not using\n", curRoom.Name)
				removedRooms[curRoom.Name] = true
			}
		}
	}
	// To Do Remove unusing paths with GetUsingRoomNamesFromTo

	fmt.Printf("Removed Roooms: %+v\n", removedRooms)
	return nil
}
