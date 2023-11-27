package main

import "fmt"

func status(name string) {
	balls := map[string]int{"김": 10, "와": 0}
	//ball := balls[name]
	//fmt.Println(ball)
	ball, exists := balls[name]

	if !exists { // true면 ㅇㅇ false면 없
		fmt.Println(name, "님은 게임 참여자가 아닙니다.")
	} else if ball < 1 {
		fmt.Println(name, "님은", balls[name], "개로 탈락했습니다")
	} else {
		fmt.Println(name, "님은 게임에서 승리하셨습니다.")
	}
}

func main() {
	status("김")
	status("와")
	status("왁")
}
