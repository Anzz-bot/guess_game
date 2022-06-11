//package main
//
//import (
//	"fmt"
//	"math/rand"
//)
//
//func main() {
//	maxNum := 100
//	secretNumber := rand.Intn(maxNum)
//	fmt.Println("The secret number is ", secretNumber)
//}
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

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured :", err)
			return
		}
		input = strings.TrimSuffix(input, "\r\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input err:", err)
		}
		fmt.Println("your guess is", guess)
		if guess != secretNumber {
			fmt.Println("you is wrong please guess again")
			continue
		} else {
			fmt.Println("you are right")
			break
		}
	}
	fmt.Println("game over")
}
