package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') // option 2, ignore  the errir return value with the
	log.Fatal(err)
	fmt.Println(inputScore)
}
