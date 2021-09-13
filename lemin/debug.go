package lemin

import "fmt"

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
