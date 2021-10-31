package anthill

// Modes for FieldInfo
const (
	FIELD_ANTS  = iota // On Reading Ants
	FIELD_ROOMS        // On Reading Rooms
	FIELD_PATHS        // On Reading Paths | Relations
)

// Modes for Path
const (
	REVERSED = -1 // directed, REVERSED path (from end to start)
	BLOCKED  = 0  // blocked path (from start to end)
	STABLE   = 1  // double directed path
)
