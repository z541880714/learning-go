package collection

import (
	"fmt"
	"slices"
)

var aa []string = nil

func SliceTest() {
	//slice1()
	//slice2()
	//slicePoint()
	//sliceCap()
	pointTest()
}

func pointTest() {
	a := new(int)
	fmt.Printf("%d, %p, type:%T\n", *a, a, a)

	i := 1
	var p *int = &i
	fmt.Println("p value:", *p)
	*p = 30
	fmt.Println("p value:", *p)
}

func slice1() {
	fmt.Println("len:", len(aa))

	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}

	var s2 []int = nil
	s3 := make([]int, 0)
	fmt.Println(s1, s2, s3)

	var s4 []int = make([]int, 0, 10)
	fmt.Println(s4)

	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	s6 = arr[1:4] // [1, 4)  左开右闭
	fmt.Println(s6)
}

func slice2() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 []int = arr[0:5]
	fmt.Println("slice0:", slice0)
	var slice1 = arr[:5]
	fmt.Println("slice1:", slice1)
	var slice2 = arr[:]
	fmt.Println("slice2:", slice2)
	slice3 := arr[:len(arr)-1]
	fmt.Println("slice3:", slice3, "cap:", cap(slice3), len(slice3))
}

func slicePoint() {
	var d [5]struct {
		x int
	}

	s := d[1:4] // s切片只是一个引用

	d[1].x = 10
	s[1].x = 20

	fmt.Println(d)
	fmt.Println(s)
	fmt.Printf("%p,  %p \n", &d, &d[0])
	fmt.Printf("%p, %p \n", &d[1], &s[0])
}

func sliceCap() {
	var arr = make([]int, 0, 1)
	c := cap(arr)

	for i := 0; i < 100; i++ {
		arr = append(arr, i)
		if n := cap(arr); n > c {
			fmt.Println("cap:", cap(arr), " len:", len(arr))
			c = n
		}
	}

	// copy slice
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("%v\n", s2)
	copy(s2[1:], s1)
	fmt.Printf("%v\n", s2)
	slices.Delete(s1, 0, 1)
	fmt.Println("delete  s1:", s1)

}
