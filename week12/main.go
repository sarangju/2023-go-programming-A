package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := make([]int, 5) // 단축 연산자( make함수를 이용해 슬라이스 생성 후 메모리 할당, 제로 값 적용)

	s := []int{0, 0, 0, 0, 0}
	// 슬라이스 리터럴을 이용해 슬라이스 생성 및 메모리 할당, 초기화 진행

	s[4] = 99
	s[2] = 91

	for _, value := range s {
		fmt.Println(value)
	}

	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}

	test := [3]string{"inha", "go", "student"} // 배열 리터럴을 이용해서test 배열 생성
	// testS := test[0:4] // invalid argument: index 4 out of bounds [0:4]
	testS := test[0:2] // testS := test[:2]
	testS2 := test[1:]
	testS2[0] = "python"
	// 원본값 어찌구...
	fmt.Println(testS2)
	fmt.Println(testS, len(testS))
	fmt.Println(test)

	// 으아아아악
	a := []string{"a", "b", "c", "d"}
	aS := a[:2]
	aS[1] = "z"
	fmt.Println(a)
	fmt.Println(aS)

	b := [4]int{4, 3, 2, 1}
	bS := b[1:3]
	fmt.Println(bS)

	c := make([]string, 4, 5) // string은 타입, 4는 length(개수), 5는 capacity(수용량)
	c[0] = "a"
	c[2] = "c"
	c[3] = "d"
	cS := c[0:2]
	cS[1] = "z"
	fmt.Println(c, len(c), cap(c))
}
