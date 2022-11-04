package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ringTime1 = time.Date(2022, 11, 4, 6, 20, 0, 0, time.Local)
	ringTime2 = time.Date(2022, 11, 4, 14, 0, 0, 0, time.Local)
	wg        = sync.WaitGroup{}
)

func main() {
	wg.Add(2)
	go clock1(ringTime1)
	go clock2(ringTime2)
	wg.Wait()
}

func ring() {
	timer := time.NewTimer(10 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-timer.C:
			return
		case <-ticker.C:
			fmt.Println("滴滴滴！")
		}
	}
}

func clock1(rt1 time.Time) {
	now := time.Now()
	for rt1.Before(now) {
		rt1 = rt1.Add(3600 * time.Second)
	}
	for {
		now = time.Now()
		subTime := rt1.Sub(now)
		ticker := time.NewTicker(subTime)
		select {
		case <-ticker.C:
			ring()
		}
		rt1 = rt1.Add(3600 * time.Second)
	}
	wg.Done()
}

func clock2(rt2 time.Time) {
	now := time.Now()
	for rt2.Before(now) {
		rt2 = rt2.Add(86400 * time.Second)
	}
	for {
		now := time.Now()
		subTime := rt2.Sub(now)
		ticker := time.NewTicker(subTime)
		select {
		case <-ticker.C:
			ring()
		}
		rt2 = rt2.Add(86400 * time.Second)
	}
	wg.Done()
}
