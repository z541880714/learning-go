package cpuPercent

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func MemInfo() {
	v, _ := mem.VirtualMemory()
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	fmt.Println(v)

}

func sum(arr []float64) float64 {
	sum := float64(0)
	len_ := len(arr)
	if len_ == 0 {
		return 0
	}

	for _, v := range arr {
		sum += float64(v)
	}
	return sum / float64(len_)
}

func getNowFmt() string {
	now := time.Now()
	return fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func CalCpuInfo() {
	group := sync.WaitGroup{}
	group.Add(60 * 60 * 24)

	f, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Print("error:", err)
		}
	}(f)

	go func() {
		for {
			// 异步进行...
			percents, _ := cpu.Percent(time.Second, true)
			ret := sum(percents)
			retStr := fmt.Sprintf("%.2f\n", ret)
			log.Println(retStr)
			timeStr := getNowFmt()
			// line := strconv.FormatFloat(sum(percents), 'e', -1, 64)
			_, err2 := f.WriteString(timeStr + " " + retStr)
			if err2 != nil {
				return
			}
			group.Done()
		}
	}()
	log.Println("函数执行完成,  等待 work 结束...")
	// 等待结束.
	group.Wait()
}
