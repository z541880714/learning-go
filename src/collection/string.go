package collection

import "fmt"

const (
	a  = iota
	b  = 100
	cc = 1
	c  = iota

	d
)

const (
	n1, n2 = iota + 100, iota + 2

	n4 = iota
	n5
	n3
)

func Array1() {
	arr := make([]int, 10)
	for index, it := range arr {
		fmt.Print(index, it)
	}

	var arr2 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("arr2:", arr2)

	arr3 := [5]int{1: 100, 4: 400}
	fmt.Println("arr3:", arr3)

	arr4 := [5]int{1, 3}
	fmt.Println("arr4:", arr4)

	arr5 := [...]int{1, 2, 3}
	fmt.Println("arr5:", arr5)

	arrS := [...]struct {
		name string
		age  int
	}{
		{"jack", 10},
		{"tom", 12},
	}

	fmt.Println("arrS:", arrS)
}

func MulDimArray() {
	var arr0 = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("arr0:", arr0)

	arr0[0][1] = 100
	fmt.Println("arr0:", arr0)

	//遍历
	for k1, v1 := range arr0 {
		for k2, v2 := range v1 {
			fmt.Println(fmt.Sprintf("%d:%d-> %d", k1, k2, v2))
		}
	}
}

func modifyArr(arr [5]int) {
	arr[0] = 100
	fmt.Println("in func arr:", arr)
}

func modifyArrPoint(arr *[5]int) {
	arr[0] = 100
	fmt.Println("in func arr:", arr)
}

// FuncParam 数组传参 拷贝与引用.
func FuncParam() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr1:", arr1)
	modifyArr(arr1)
	fmt.Println("333", arr1)
	modifyArrPoint(&arr1)
	fmt.Println("444", arr1)
}
