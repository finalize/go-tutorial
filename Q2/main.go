package main

import "fmt"

func main() {
	// Q1
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	i := l[0]
	for _, v := range l {
		if i > v {
			i = v
		}
	}
	fmt.Println(i)

	// Q2
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	var uint int
	for _, v := range m {
		uint += v
	}
	fmt.Println(uint)
}
