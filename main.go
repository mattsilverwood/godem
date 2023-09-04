// Program godem
package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Replace "yourfile.bin" with the path to your binary file
	filePath := "test.dem"

	// Read the binary file into a []byte slice
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var dem DemoHeader
	binary.Read(file, binary.LittleEndian, &dem)

	if strings.Compare(string(dem.Header[:]), "HL2DEMO") != 0 {
		panic("Not a source demo file.")
	}

	// fmt.Printf("Found Demo file in %s\n", file.Name())
	// fmt.Printf("HEADER: \t\t%s\n", dem.Header)
	// fmt.Printf("DEMO PROTOCOL VER: \t%d\n", dem.DemoProtocol)
	// fmt.Printf("NETWORK PROTOCOL: \t%d\n", dem.NetworkProtocol)
	// fmt.Printf("SERVER NAME: \t\t%s\n", dem.ServerName[:])
	// fmt.Printf("CLIENT NAME: \t\t%s\n", dem.ClientName[:])
	// fmt.Printf("MAP NAME: \t\t%s\n", dem.MapName)
	// fmt.Printf("GAME: \t\t\t%s\n", dem.GameDirectory)
	// fmt.Printf("GAME DIRECTORY: \t%s\n", dem.GameDirectory)
	// fmt.Printf("PLAYBACK TIME (sec): \t%f\n", dem.PlaybackTime)
	// fmt.Printf("TICKS: \t\t\t%d\n", dem.Ticks)
	// fmt.Printf("FRAMES: \t\t%d\n", dem.Frames)
	// fmt.Printf("SIGNONLENGTH: \t\t%d\n", dem.SignOnLength)
}
