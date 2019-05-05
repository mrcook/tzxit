package tzx

import (
	"fmt"

	"github.com/mrcook/tzxit/tape"
)

// GlueBlock
// ID: 5Ah (90d)
// This block is generated when you merge two ZX Tape files together. It is here so that you can
// easily copy the files together and use them. Of course, this means that resulting file would
// be 10 bytes longer than if this block was not used. All you have to do if you encounter this
// block ID is to skip next 9 bytes.
// If you can avoid using this block for this purpose, then do so; it is preferable to use a
// utility to join the two files and ensure that they are both of the higher version number.
type GlueBlock struct {
	Value [9]byte // BYTE[9] Value: { "XTape!",0x1A,MajR,MinR } Just skip these 9 bytes and you will end up on the next ID.
}

func (g *GlueBlock) Read(file *tape.File) {
	for i, b := range file.ReadBytes(9) {
		g.Value[i] = b
	}
}

func (g GlueBlock) Id() uint8 {
	return 0x5a
}

func (g GlueBlock) Name() string {
	return "Glue Block"
}

// ToString returns a human readable string of the block data
func (g GlueBlock) ToString() string {
	return fmt.Sprintf("> %s", g.Name())
}
