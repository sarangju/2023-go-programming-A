package main

import "fmt"

func main() {
	/* var games map[int]string
	games = make(map[int]string) */

	games := make(map[int]string)

	// append
	games[456] = "우우우"
	games[722] = "최최최"
	games[303] = "정정정"
	games[125] = "김김김"
	games[912] = "유유유"
	games[203] = "왕왕왕"

	for _, v := range games {
		fmt.Println(v)
	}

	games[456] = "곡곡곡" // update
	delete(games, 912) // delete

	for k, v := range games {
		fmt.Println(k, v)
	}
}
