package main

import 	"fmt"


func main() {
	var i int

	for i = 0; i < 10; i++ {
		if i%10 == 0 {
			fmt.Printf("%d\t", i)
		}
	}
	for i = 10; i < 100; i++ {
		if i%10 == 0 || (i%10+i/10)%10 == 0 {
			fmt.Printf("%d\t", i)
		}
	}
	for i = 100; i < 1000; i++ {
		if i%100 == 0 {
			fmt.Printf("%d\t", i)
		} else if i%10 == 0 && (i/100+i%100/10)%10 == 0 {
			fmt.Printf("%d\t", i)
		} else if i%10 != 0 && (i-i%10)%100 == 0 && (i/100+i%10)%10 == 0 {
			fmt.Printf("%d\t", i)
		} else {
			continue
		}
	}
	for i = 1000; i < 10000; i++ {
		if i%1000 == 0 {
			fmt.Printf("%d\t", i)
		} else if i%100 == 0 && (i/1000+i%1000/100)%10 == 0 {
			fmt.Printf("%d\t", i)
		} else if i%10 == 0 && i%100 != 0 && (i-i%100)%1000 == 0 && (i/1000+i%100/10)%10 == 0 {
			fmt.Printf("%d\t", i)
		} else if i%10 != 0 && (i-i%10)%1000 == 0 && (i/1000+i%10)%10 == 0 {
			fmt.Printf("%d\t", i)
		} else {
			continue
		}
	}
}
