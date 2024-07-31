package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 99.9
	b := 99.9
	var bb = strconv.FormatFloat(b, 'e', 10, 64)
	var aa = strconv.FormatFloat(a, 'e', 10, 32)
	fmt.Println("bb:", bb, "aa:", aa)
	//cpuPercent.CalCpuInfo()
}
