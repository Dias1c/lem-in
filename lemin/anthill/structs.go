package anthill

type anthill struct {
	AntsCount  int
	StepsCount int
	Start, End string
	Rooms      map[string]*room
	Paths      []*list
	Result     []string
}

type room struct {
	Name                string           // Name
	X, Y                int              // Coordinates
	PathsIn, PathsOut   map[string]*room // Using For Surballe algo
	MarkedIn, MarkedOut bool             // For know using rooms
	ParentIn, ParentOut *room            // Using For Surballe algo
	UsingOnPath         bool             // Is Using On Finded Paths
}

type node struct {
	Room *room
	Next *node
}

type list struct {
	Len   int
	Front *node
	Back  *node
}
