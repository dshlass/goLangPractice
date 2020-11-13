package main

import "fmt"

func main() {
	//41 90
	x := 1
	for {
		fmt.Println(x)
		if x >= 10000 {
			break
		}
		x++
	}
}
