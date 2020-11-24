package binary

import "encoding/binary"

// Copy predefined byte orders from standard library.
var (
	BigEndian    = binary.BigEndian
	LittleEndian = binary.LittleEndian
)
