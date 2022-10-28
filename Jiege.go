package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup

func main() {
	wg.Add(1)
	好康的 := 欢迎来我家玩1(打电动)
	go 欢迎来我家玩2()
	wg.Wait()
	fmt.Println(好康的)
}

func 欢迎来我家玩1(d func() string) string {
	time.Sleep(2 * time.Second)
	return d()
}

func 欢迎来我家玩2() {
	fmt.Println("登dua郎")
	wg.Done()
}

func 打电动() string {
	return "输了啦，都是你害的"
}
