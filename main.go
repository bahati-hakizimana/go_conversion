package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("welcome to our peezer and conversion app")
    fmt.Println("please rate our peezer between 1 and 5")

	reader :=bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating number:", numRating+1)
	}
}