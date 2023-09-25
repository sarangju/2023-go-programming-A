package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // get the current data and time as an integer
	answer := rand.Intn(100) + 1 // random integer num between 1-100
	fmt.Println("Guess Number Game")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)

	for i :=0; i< 10; i++{
		fmt.Println("You have ", 10-i, "chances")
		fmt.Print("Input guess number: ")
		inputNumber, err := reader.ReadString(`\n`)
		if err != nill {
			log.Fatal(err)
		}
		inputNumberString = strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoil(inputNumberString)
		if err != nill {
			log.Fatal(err)
		}
		if inputNumber < answer {
			fmt.Println("Guess number is lower than answer") // answer is higher than inputnum
		} else if inputNumber > answer {
			fmt.Println("Guess number is higher than answer")
		}
	}|
