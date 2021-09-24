package anthill

import (
	"container/list"
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

func (r *room) PrintRoomInfo() {
	if r == nil {
		fmt.Println("Room is Nil")
		return
	}
	fmt.Printf("Room: %q, Pos: %d, %d;\nRooms: %+v\n", r.Name, r.X, r.Y, r.Paths)
}

func PrintRoomsInLinkedList(l *list.List) {
	fmt.Println("PrintRoomsInLinkedList:")
	for n := l.Front(); n != nil; n = n.Next() {
		t := n.Value.(*room)
		fmt.Printf("%v", t.Name)
		if n.Next() != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
