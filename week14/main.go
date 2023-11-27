package main

import "fmt"

func main() {
	
	games:= map[int]string {
		456:"우우우",
		722:"최최최",
		303:"정정정",
		125:"김김김",
		912:"유유유",
		203:"왕왕왕," // 마지막에 쉼표 찍어줘야 ㅇㅇ
	}

	for _, v := range games {
		fmt.Println(v)
	}

	games[456] = "곡곡곡" // update
	delete(games, 912) // delete

	for k, v := range games {
		fmt.Println(k, v)
	}
}
