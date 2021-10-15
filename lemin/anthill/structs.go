package anthill

// The found paths are saved in Result. Using for save write result to writer
type Result struct {
	AntsCount int
	Paths     []*list
}

// Stores information about the graph, the data being read, and the result. Using for find paths
type anthill struct {
	// For Reading Data
	FieldInfo *fieldInfo
	// Main Data
	AntsCount  int
	Start, End string
	Rooms      map[string]*room
	// Results
	StepsCount int
	Result     *Result
}

type room struct {
	Name                string           // Name
	X, Y                int              // Coordinates
	PathsIn, PathsOut   map[string]*room // Using For Surballe algo
	MarkedIn, MarkedOut bool             // For know using rooms
	ParentIn, ParentOut *room            // Using For Surballe algo
	UsingOnPath         bool             // Means - Is Using On Finded Paths
}

// with fieldInfo, we understand What data we fill in for the anthill
type fieldInfo struct {
	MODE             byte                 // FIELD_ANTS | FIELD_ROOMS | FIELD_PATHS
	Start, End       bool                 // Should Be True
	IsStart, IsEnd   bool                 // For Know Which Room is Reading
	UsingCoordinates map[int]map[int]bool // Chekking for unique Coordinates on Rooms
}

// List of node wich has room.
type node struct {
	Room *room
	Next *node
}

// List of Room nodes. Used to store found paths
type list struct {
	Len   int
	Front *node
	Back  *node
}

type antStruct struct {
	Num  int
	Path int
	Pos  int
	Next *antStruct
}

// queue of ants. Used for write ant position on every step
type antQueue struct {
	Front *antStruct
	Back  *antStruct
}
