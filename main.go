package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	script := []byte(`loadstring(game:HttpGet("https://raw.githubusercontent.com/EdgeIY/infiniteyield/master/source"))()` + "\x00")
	header := make([]byte, 16)
	binary.LittleEndian.PutUint32(header[8:12], uint32(len(script)))

	if conn, err := net.DialTimeout("tcp", "127.0.0.1:5553", 3*time.Second); err == nil {
		defer conn.Close()
		conn.Write(append(header, script...))
		fmt.Println("F12 in Roblox to see script activity.")
	} else {
		fmt.Println("Error:", err)
	}
}
