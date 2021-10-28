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
	Name                string        // Name
	X, Y                int           // Coordinates
	Paths               map[*room]int // state -> REVERSED || BLOCKED || STABLE
	ParentIn, ParentOut *room         // Using for Surballe algo
	VisitIn, VisitOut   bool
	Weight              [2]int // Out weight in 0 index, In in 1
	Separated           bool   // Flag for checking separated node
	InNewPath           bool   // Flag for using in replace edges
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

// for sorting rooms
type weightNode struct {
	Room   *room
	Weight int
	Mark   bool //false if it's in, true if it's out
	Next   *weightNode
}

// queue for sorted rooms by weight
type sortedQueue struct {
	Front *weightNode
	Back  *weightNode
}
