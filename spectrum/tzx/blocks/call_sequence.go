package blocks

import (
	"fmt"

	"retroio/storage"
)

// CallSequence
// ID: 26h (38d)
// This block is an analogue of the CALL Subroutine statement. It basically executes a sequence of
// blocks that are somewhere else and then goes back to the next block. Because more than one call
// can be normally used you can include a list of sequences to be called. The 'nesting' of call
// blocks is also not allowed for the simplicity reasons. You can, of course, use the CALL blocks
// in the LOOP sequences and vice versa. The value is relative for the obvious reasons - so that
// you can add some blocks in the beginning of the file without disturbing the call values. Please
// take a look at 'Jump To Block' for reference on the values.
type CallSequence struct {
	Count  uint16   // N WORD  Number of calls to be made
	Blocks []uint16 // WORD[N] Array of call block numbers (relative-signed offsets)
}

// Read the tape and extract the data.
// It is expected that the tape pointer is at the correct position for reading.
func (c *CallSequence) Read(reader *storage.Reader) {
	c.Count = reader.ReadShort()

	for i := 0; i < int(c.Count); i++ {
		c.Blocks = append(c.Blocks, reader.ReadShort())
	}
}

// Id of the block as given in the TZX specification, written as a hexadecimal number.
func (c CallSequence) Id() uint8 {
	return 0x26
}

// Name of the block as given in the TZX specification.
func (c CallSequence) Name() string {
	return "Call Sequence"
}

// ToString returns a human readable string of the block data
func (c CallSequence) ToString() string {
	str := fmt.Sprintf("%s\n", c.Name())
	for _, b := range c.Blocks {
		str += fmt.Sprintf(" - %d\n", b)
	}
	return str
}

// ReturnFromSequence
// ID: 27h (39d)
// This block indicates the end of the Called Sequence. The next block played will be the block after
// the last CALL block (or the next Call, if the Call block had multiple calls).
// This block has no body.
type ReturnFromSequence struct{}

// Read the tape and extract the data.
// It is expected that the tape pointer is at the correct position for reading.
func (r ReturnFromSequence) Read(reader *storage.Reader) {}

// Id of the block as given in the TZX specification, written as a hexadecimal number.
func (r ReturnFromSequence) Id() uint8 {
	return 0x27
}

// Name of the block as given in the TZX specification.
func (r ReturnFromSequence) Name() string {
	return "Return from Sequence"
}

// ToString returns a human readable string of the block data
func (r ReturnFromSequence) ToString() string {
	return r.Name()
}