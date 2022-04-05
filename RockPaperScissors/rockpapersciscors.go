package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Ready program!
func main() {
	rand.Seed(time.Now().UnixNano())
	var userChoice string
	var userChoiceIndex int
	var machineChoiceInt int
	fmt.Println("What do you choose? (rock, paper or scissors)")
	fmt.Scanln(&userChoice)
	machineChoiceInt = rand.Intn(3)
	fmt.Println(machineChoiceInt)

	switch {
	case userChoice == "rock":
		userChoiceIndex = 0
	case userChoice == "paper":
		userChoiceIndex = 1
	case userChoice == "scissors":
		userChoiceIndex = 2
	}

	if userChoiceIndex == machineChoiceInt {
		fmt.Println("Draw!")
	} else {
		if userChoiceIndex == 0 && machineChoiceInt == 1 {
			fmt.Println("You lost!")
		} else {
			if userChoiceIndex == 0 && machineChoiceInt == 2 {
				fmt.Println("You won!")
			} else {
				if userChoiceIndex == 1 && machineChoiceInt == 0 {
					fmt.Println("You won!")
				} else {
					if userChoiceIndex == 1 && machineChoiceInt == 2 {
						fmt.Println("You lost!")
					} else {
						if userChoiceIndex == 2 && machineChoiceInt == 0 {
							fmt.Println("You lost!")
						} else {
							if userChoiceIndex == 2 && machineChoiceInt == 1 {
								fmt.Println("You won!")
							}
						}
					}
				}
			}
		}

	}
	fmt.Println("Press ENTER to quit program.....")
	fmt.Scanln()
}
