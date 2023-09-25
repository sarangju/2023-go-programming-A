package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') // option 2
	if err != nil {
		log.Fatal(err)
	}
	inputScoreString = strings.TrimSpace(inputScoreString)      // remove space bar
	inputScore, err := strconv.ParseFloat(inputScoreString, 32) // casting
	var grade string
	if inputScore >= 90 {
		grade := "A grade!"
	} else {
		grade := "under A grade~"
	}
	fmt.Println("You got " + grade) // 왜 에러남 자꾸
}
