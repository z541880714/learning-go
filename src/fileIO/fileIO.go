package fileIO

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bytesToStrings(bytes []byte) []string {
	var _len = len(bytes)
	var ret = make([]string, _len)
	for i, it := range bytes {
		fmt.Println("index:", i, "item:", string(it))
		str := string(it)
		ret = append(ret, str)
	}
	return ret
}

func ReadFile() string {
	fmt.Println("start read file ")
	f, _ := os.Open("log.txt")
	defer f.Close()
	var buffer = bufio.NewReader(f)
	var contentList = make([]string, 0)
	for {
		lineBytes, _, err := buffer.ReadLine()
		if err != nil {
			break
		}
		contentList = append(contentList, string(lineBytes))
	}
	ret := strings.Join(contentList, "\n")
	return ret
}

func WriteFile(content string, path string) {

}
