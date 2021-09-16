package lemin

import (
	"fmt"
)

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
		return nil, fmt.Errorf("GetUsingRoomNamesFromTo: RoomFrom: %s (%+v) or RoomTo: $s (%+v) Not Found!", from, t.Rooms[from], to, t.Rooms[to])
	}
	countRooms := len(t.Rooms)
	moved := make(map[string]bool, countRooms)
	moved[from] = true
	moved[to] = true
	// Start Match
	stack := []*room{t.Rooms[from]}
	remains := 1
	for remains >= 0 {
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

func (t *terrain) RemoveUselessFromLevels(levels []map[string]*room) {
	countLevels := len(levels)
	// Forriben
	// forbidden := make(map[string]bool, countLevels)
	// forbidden[t.Start] = true
	// forbidden[t.End] = true
	// Set Forribens
	// for _, level := range levels {
	// 	for name := range level {
	// 		forbidden[name] = true
	// 	}
	// }

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
			}
		}
	}
	// To Do Remove unusing paths with GetUsingRoomNamesFromTo

	// Romoves only 1
	// for lvlIdx, level := range levels {
	// 	//Set forriben
	// 	for name := range level {
	// 		forbidden[name] = true
	// 	}
	// 	// Remove rooms from level
	// 	for _, curRoom := range level {
	// 		if len(curRoom.Paths) == 1 { // Remove if has only 1 path
	// 			delete(levels[lvlIdx], curRoom.Name)
	// 			removedRooms[curRoom.Name] = true
	// 			countRooms--
	// 		} else { // Remove if has only bad paths
	// 			isHavePaths := false
	// 			for name := range curRoom.Paths {
	// 				if !forbidden[name] {
	// 					isHavePaths = true
	// 					break
	// 				}
	// 			}
	// 			if !isHavePaths {
	// 				delete(levels[lvlIdx], curRoom.Name)
	// 				removedRooms[curRoom.Name] = true
	// 				countRooms--
	// 			}
	// 		}
	// 	}
	// }
	//
	fmt.Printf("Removed Roooms: %q\n", removedRooms)
}
