package fileIO

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

// init 函数自动优先执行... 不需要被调用..
func init() {
	fmt.Println("this is a init method .....")
}

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
