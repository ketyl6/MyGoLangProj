package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//ASCII nums 048 to 122

func main() {
	rand.Seed(time.Now().UnixNano())
	asciiCodes := [13]int{}
	for i := 0; i <= 12; {
		asciiCodes[i] = rand.Intn(74) + 48
		i++
	}
	fmt.Println("This is your new 13 character randomly generated password: ")
	for z := 0; z <= 12; z++ {
		fmt.Printf("%c", asciiCodes[z])
	}
	fmt.Print("\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
