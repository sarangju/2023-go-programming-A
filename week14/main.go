package main

import "fmt"

func main() {
	//   var balls map[string]int
	balls := make(map[string]int)
	fmt.Printf("%#v\n", balls)
	balls["최최최"] = 777
	fmt.Println(balls{"최최최"})
	fmt.Println(balls{"없는사람"})
}
