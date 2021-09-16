package lemin

import "fmt"

// This File for DEBUG

// PrintTerrainDatas - Shows terrain Data
func PrintTerrainDatas(terrain *terrain) {
	if terrain == nil {
		fmt.Println("PrintTerrainDatas: Terrain is Null")
	}
	fmt.Println("+---: Name\tX:\tY:")
	for _, room := range terrain.Rooms {
		fmt.Printf("Room: %v  \t%d\t%d\n", room.Name, room.X, room.Y)
	}
	fmt.Printf("StartRoom: %v\nEndRoom: %v\nCountAnts: %v\n", terrain.Start, terrain.End, terrain.AntsCount)
}

// PrintLevels - prints levels
func PrintLevels(levels []map[string]*room) {
	fmt.Println("+==== Print Levels ====+")
	for level, rooms := range levels {
		fmt.Printf("Level: %d\n", level)
		for name := range rooms {
			fmt.Printf("%q ", name)
		}
		fmt.Println()
	}
	fmt.Println("+=====\t\t=======+")
}

func (r *room) PrintRoomInfo() {
	fmt.Printf("Room: %q, Pos: %d, %d;\nRooms: %+v\n", r.Name, r.X, r.Y, r.Paths)
}
