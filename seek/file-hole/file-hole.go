package main

import (
	"fmt"
	"os"
)

func main() {
	// fd, err := unix.Open("file-hole.txt", unix.O_CREAT|unix.O_WRONLY|unix.O_TRUNC, unix.O_RDWR)
	// if err != nil {
	// 	fmt.Println("eror", err)
	// }
	// if fd < 0 {
	// 	fmt.Println("System error")
	// }
	// os.lseek()
	file, err := os.Create("test1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Seek(5, 0)
	byteSlice := []byte("This is a testing of a hole in a g g gfile!\n")
	file.Write(byteSlice)
	file.Seek(150, 2)
	byteSlice2 := []byte("Bytes!\n")
	file.Write(byteSlice2)
}
