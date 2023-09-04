package main

// DemoHeader structures a .DEM file header.
// https://developer.valvesoftware.com/wiki/DEM_(file_format)#Demo_header
type DemoHeader struct {
	Header          [8]byte   // 8 characters, should be "HL2DEMO"+NULL
	DemoProtocol    int32     // Demo protocol version (stored in little endian)
	NetworkProtocol int32     // Network protocol version number (stored in little endian)
	ServerName      [260]byte // 260 characters long
	ClientName      [260]byte // 260 characters long
	MapName         [260]byte // 260 characters long
	GameDirectory   [260]byte // 260 characters long
	PlaybackTime    float32   // The length of the demo, in seconds
	Ticks           int32     // The number of ticks in the demo
	Frames          int32     // The number of frames in the demo
	SignOnLength    int32     // Length of the signon data (Init for first frame)
}
