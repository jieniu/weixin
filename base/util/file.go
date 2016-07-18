package util

import (
	"fmt"
	"os"
)

// ReadFile read small file content which less than 1024
func ReadFile(filename string) (buf []byte, err error) {
	fin, err := os.Open(filename)
	defer fin.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	buf = make([]byte, 1024)
	var len int
	for {
		n, _ := fin.Read(buf)
		len += n
		if 0 == n {
			break
		}
	}
	return buf[:len], nil
}
