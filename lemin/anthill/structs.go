package anthill

type Result struct {
	AntsCount int
	Paths     []*list
}

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

type fieldInfo struct {
	MODE             byte                 // FIELD_ANTS | FIELD_ROOMS | FIELD_PATHS
	Start, End       bool                 // Should Be True
	IsStart, IsEnd   bool                 // For Know Which Room is Reading
	UsingCoordinates map[int]map[int]bool // Chekking for unique Coordinates on Rooms
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

type antStruct struct {
	Num  int
	Path int
	Pos  int
	Next *antStruct
}

type antQueue struct {
	Front *antStruct
	Back  *antStruct
}

type antList struct {
	Front *antStruct
	Back  *antStruct
}
