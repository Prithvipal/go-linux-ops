package main

import (
	"fmt"
	"io"

	"golang.org/x/sys/unix"
)

func main() {
	if newOffset, _ := unix.Seek(unix.Stdin, 0, io.SeekCurrent); newOffset == -1 {
		fmt.Println("Can not seek")
	} else {
		fmt.Println("Seek OK!!!")
	}
}
