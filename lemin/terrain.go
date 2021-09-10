package lemin

type terrain struct {
	AntsCount  int              // Count Ants in Terrain
	Start, End string           // Start, End Room Names
	Rooms      map[string]*room // map["RoomName"]*Room
}

func getTerrainFromLines(lines []string) (*terrain, error) {
	//
	countAnts, err := getCountAntsFromLine(&lines)
	if err != nil {
		return nil, err
	}
	rooms, startRoom, endRoom, err := getRoomsFromLines(&lines)
	if err != nil {
		return nil, err
	}
	err = setPathsFromLines(&lines, rooms)
	if err != nil {
		return nil, err
	}
	return &terrain{AntsCount: countAnts, Rooms: rooms, Start: startRoom, End: endRoom}, nil
}

func (*terrain) Match() (string, error) {
	return "", nil
}
