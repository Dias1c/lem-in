package anthill

import (
	"fmt"
)

// This File for DEBUG
// PrintTerrainDatas - Shows anthill Data
func PrintTerrainDatas(anthill *anthill) {
	if anthill == nil {
		fmt.Println("PrintTerrainDatas: Terrain is Null")
		return
	}
	fmt.Println("+---: Name\tX:\tY:")
	for _, room := range anthill.Rooms {
		fmt.Printf("Room: %v  \t%d\t%d\n", room.Name, room.X, room.Y)
	}
	fmt.Printf("StartRoom: %v\nEndRoom: %v\nCountAnts: %v\n", anthill.Start, anthill.End, anthill.AntsCount)
}

func printPaths(paths []*list) {
	fmt.Println("###")
	for _, l := range paths {
		start, end := l.Front, l.Back
		for start != end {
			fmt.Printf("%s --> ", start.Room.Name)
			start = start.Next
		}
		fmt.Printf("%s || len=%d\n", end.Room.Name, l.Len)
	}
	fmt.Println("###")
}
