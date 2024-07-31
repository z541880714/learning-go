package fileIO

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestWriteFile(t *testing.T) {
	f, _ := os.Create("test.txt")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("error:", err)
		}
	}(f)

	for i := 0; i < 10; i++ {
		f.WriteString(strconv.Itoa(i) + "\n")
	}
}

func TestReadFile(t *testing.T) {
	ret := ReadFile()
	fmt.Println("ret:\n" + ret)
}
