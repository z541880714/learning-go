package collection

import (
	"encoding/json"
	"fmt"
	"time"
)

func mapTest() {
	//test1()
	//test2()
	test3()
}

func test3() {
	var a, b, c int
	fmt.Println(a, b, c)

	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received: %d from c1\n", i1)
	case i2 = <-c2:
		fmt.Printf("receved: %d from c2\n", i2)
	case i3, ok := <-c3:
		if ok {
			fmt.Printf("received: %d from c3\n", i3)
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication \n")
	}

	var resChan = make(chan int)
	select {
	case data := <-resChan:
		func(data int) {
			fmt.Println("case 1 invoked ...")
		}(data)
	case t := <-time.After(time.Second * 3):
		fmt.Println("time out.... t:", t)

	}

	println("--------------------------")

}

type Class struct {
	name string
}

type Student struct {
	ID     int
	Gender string
	name   string
	class  *Class
}

func test2() {
	s1 := &Student{
		ID:     0,
		Gender: "man",
		name:   "jack",
		class: &Class{
			name: "one",
		},
	}
	fmt.Printf("%v\n", s1)
	data, _ := json.Marshal(s1)
	fmt.Printf("data: %s\n", data)

}
func test1() {
	fmt.Println("map test start ....")
	var m1 = map[string]int{"1": 10, "2": 20, "3": 30}
	fmt.Println("m1:", m1)

	var m2 = make(map[string]int, 0)
	m2["1"] = 1
	m2["2"] = 2
	fmt.Println("m2:", m2)

	delete(m2, "1")
	fmt.Println("m2:", m2)

	value, ok := m2["2"]
	fmt.Println("value:", value, "ok:", ok)

	value, ok = m2["3"]
	fmt.Println("value:", value, " ok:", ok)

	//遍历..
	for k, v := range m1 {
		fmt.Println("k:", k, " v:", v)
	}
}
