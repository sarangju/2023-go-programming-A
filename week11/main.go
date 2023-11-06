package main

import (
	"fmt"
)

func main() {
	/* var primes [3]int
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	fmt.Println(primes, primes[1]) */

	/* var primes [3]int = [3]int{2, 3, 5} // primes 배열을 배열 리터럴로 초기화
	fmt.Println(primes, primes[1]) */

	primes := [3]int{2, 3, 5} // 단축 연산자
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}
	fmt.Println(test, test[3]) // boolean 타입의 제로 값

	// fmt.Println(primes[4]) // invalid argument: index 4 out of bounds [0:3]

	/* j := 0
	for j < 4 { // panic error!!!!!!
		fmt.Println(primes[j])
		j++
	} */

	i := 0
	for i < len(primes) { // while문 처럼
		fmt.Println(primes[i])
		i++
	}

	// for j := 0; j < len(primes); j++   => for문으로

	for idx, primes := range primes { // 0 2 , 1 3, 2 5
		fmt.Println(idx, primes)
	}

	// for primes := range primes { // range는 변수를 하나만 쓰면 index만 받음 => 0,1,2
	// for idx, prime := range primes { // 컴파일에러, idx 사용해야함}
	for _, primes := range primes {
		fmt.Println(primes)
	}
}
