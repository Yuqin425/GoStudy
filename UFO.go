package main

import "fmt"

func judge(name string) int {
	sum := 1
	for i := 0; i < len(name); i++ {
		for _, v := range name {
			switch v {
			case 'A':
				sum *= 1
			case 'B':
				sum *= 2
			case 'C':
				sum *= 3
			case 'D':
				sum *= 4
			case 'E':
				sum *= 5
			case 'F':
				sum *= 6
			case 'G':
				sum *= 7
			case 'H':
				sum *= 8
			case 'I':
				sum *= 9
			case 'J':
				sum *= 10
			case 'K':
				sum *= 11
			case 'L':
				sum *= 12
			case 'M':
				sum *= 13
			case 'N':
				sum *= 14
			case 'O':
				sum *= 15
			case 'P':
				sum *= 16
			case 'Q':
				sum *= 17
			case 'R':
				sum *= 18
			case 'S':
				sum *= 19
			case 'T':
				sum *= 20
			case 'U':
				sum *= 21
			case 'V':
				sum *= 22
			case 'W':
				sum *= 23
			case 'X':
				sum *= 24
			case 'Y':
				sum *= 25
			case 'Z':
				sum *= 26
			}
		}
	}
	return sum
}

func main() {
	var (
		UFO   string
		Group string
	)
	fmt.Println("请输入UFO的名字(大写字母)：")
	_, _ = fmt.Scanln(&UFO)
	fmt.Println("请输入YUAN小组的名字(大写字母)：")
	_, _ = fmt.Scanln(&Group)
	num1 := judge(UFO)
	num2 := judge(Group)
	if num1%47 == num2%47 {
		fmt.Println("GO")
	} else {
		fmt.Println("STAY")
	}
}
