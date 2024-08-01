package collection

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	// 遍历字符串
	func() {
		s := "博客可可可可可可可可可可可可可可可"
		fmt.Println("len:", len([]rune(s)))
		fmt.Println("len:", bytes.Count([]byte(s), nil))
		fmt.Println("len2:", strings.Count(s, ""))
		for i := 0; i < len(s); i++ { //byte
			fmt.Printf("%v(%c) ", s[i], s[i])
		}
		fmt.Println()
		for _, r := range s { //rune
			fmt.Printf("%v(%c) ", r, r)
		}
		fmt.Println()
	}()
}

// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
func TestModify(t *testing.T) {
	s := "hello"
	sBytes := []byte(s)
	sBytes[0] = 'a'
	fmt.Println("data:", string(sBytes))

	s = "你好中国"
	sRune := []rune(s)
	sRune[0] = 'a'
	fmt.Println("ret:", string(sRune))
}

func TestArray1(t *testing.T) {
	Array1()
}

func TestMulDimArray(t *testing.T) {
	MulDimArray()
}

func TestFuncParam(t *testing.T) {
	FuncParam()
}
