package main

import "fmt"

func main() {
	//41 90
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j <= 2; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
