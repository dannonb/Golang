package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func playGame(level int) bool {
	max := 25
	min := 1
	rand.Seed(time.Now().UnixNano())

	num1 := rand.Intn((max-min+1)+min)%level + level
	num2 := rand.Intn((max-min+1)+min)%level + level
	num3 := rand.Intn((max-min+1)+min)%level + level
	numSum := num1 + num2 + num3
	numProd := num1 * num2 * num3

	// game rules
	fmt.Println("There are three numbers in the code")
	fmt.Printf("Together, they add up to equal %d, and multiply to equal %d\n", numSum, numProd)

	//takes in user input and adds and multiplies it
	reader := bufio.NewReader(os.Stdin)

	Guessone, _ := reader.ReadString('\n')
	Guessone = strings.TrimSuffix(Guessone, "\n")

	Guesstwo, _ := reader.ReadString('\n')
	Guesstwo = strings.TrimSuffix(Guesstwo, "\n")

	Guessthree, _ := reader.ReadString('\n')
	Guessthree = strings.TrimSuffix(Guessthree, "\n")

	one, _ := strconv.Atoi(Guessone)
	two, _ := strconv.Atoi(Guesstwo)
	three, _ := strconv.Atoi(Guessthree)

	GuessSum := one + two + three
	GuessProd := one * two * three

	answer := false

	if GuessSum == numSum && GuessProd == numProd {
		fmt.Println("you are correct")
		answer = true
	} else if GuessSum != numSum || GuessProd != numProd {
		fmt.Println("you have made a mistake")
		answer = false
	}

	return answer

}

func main() {
	level := 1
	maxLevel := 5

	for level <= maxLevel {
		fmt.Printf("you are on level %d\n", level)
		complete := playGame(level)

		if complete {
			level++
		}
	}
	fmt.Println("You have beaten the game")
}
