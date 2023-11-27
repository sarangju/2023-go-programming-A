package main

import "fmt"

func main() {
	// map literal
	games := map[int]string{
		456: "우우우",
		722: "최최최",
		303: "정정정",
		125: "김김김",
		912: "유유유",
		203: "왕왕왕", // 마지막에 쉼표 찍어줘야 ㅇㅇ
	}

	name, ok := games[722]
	fmt.Println(name, ok)

	for k, v := range games {
		fmt.Println(k, v)
	}
}
