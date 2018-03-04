// Package contains definitions for all the TZX data blocks
package tape

type Block interface {
	Id() int
	Name() string
	Metadata() string
}

// Block ID table, as `HEX (DECIMAL)` values
//
// ID 10 (16): StandardSpeedData
// ID 11 (17): TurboSpeedData
// ID 12 (18): PureTone
// ID 13 (19): SequenceOfPulses
// ID 14 (20): PureData
// ID 15 (21): DirectRecording
// ID 18 (24): CswRecording
// ID 19 (25): GeneralizedData
// ID 20 (32): PauseTheTapeCommand
// ID 21 (33): GroupStart
// ID 22 (34): GroupEnd
// ID 23 (35): JumpTo
// ID 24 (36): LoopStart
// ID 25 (37): LoopEnd
// ID 26 (38): CallSequence
// ID 27 (39): ReturnFromSequence
// ID 28 (40): Select
// ID 2A (42): StopTheTapeIfIn48kMode
// ID 2B (43): SetSignalLevel
// ID 30 (48): TextDescription
// ID 31 (49): Message
// ID 32 (50): ArchiveInfo
// ID 33 (51): HardwareType
// ID 35 (53): CustomInfo
// ID 5A (90): GlueBlock

// HardwareType
// ID: 33h (51d)
// This blocks contains information about the hardware that the programs on this tape use.
// Please include only machines and hardware for which you are 100% sure that it either runs
// (or doesn't run) on or with, or you know it uses (or doesn't use) the hardware or special
// features of that machine.
type HardwareType struct {
	TypeCount uint8          // N BYTE     Number of machines and hardware types for which info is supplied
	Machines  []HardwareInfo // HWINFO[N]  List of machines and hardware
}

// HardwareInfo
// A list of hardware types and ID and sub-types can be found in the TZX specification
// (https://www.worldofspectrum.org/TZXformat.html)
type HardwareInfo struct {
	Type        uint8 // BYTE  Hardware type
	Id          uint8 // BYTE  Hardware ID
	Information uint8 // BYTE  Hardware information:
	//                           00 - The tape RUNS on this machine or with this hardware,
	//                                but may or may not use the hardware or special features of the machine.
	//                           01 - The tape USES the hardware or special features of the machine,
	//                                such as extra memory or a sound chip.
	//                           02 - The tape RUNS but it DOESN'T use the hardware
	//                                or special features of the machine.
	//                           03 - The tape DOESN'T RUN on this machine or with this hardware.
}

// CustomInfo
// ID: 35h (53d)
// This block can be used to save any information you want. For example, it might contain some
// information written by a utility, extra settings required by a particular emulator, or even
// poke data.
type CustomInfo struct {
	Identification []byte  // CHAR[10]  Identification string (in ASCII)
	Length         uint32  // L DWORD   Length of the custom info
	Info           []uint8 // BYTE[L]   Custom info
}

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
