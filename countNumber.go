package main

import "fmt"

func main() {
	var (
		n        int
		i        int
		num      int
		numCount = make(map[int]int)
	)
	fmt.Scanf("%d\n", &n)
	number := make([]int, n)
	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &number[i])
		fmt.Scanln()
	}
	for _, num = range number {
		numCount[num] += 1
	}
	for k, v := range numCount {
		fmt.Println(k, v)
	}
}
