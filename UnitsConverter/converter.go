package main

//In progress
import (
	"fmt"
)

func Weight() {
	var weightUnitfirst int
	var weightUnitsecond int
	var weightValue float32
	fmt.Println("Choose your weight unit: (1 = kilograms, 2 = pounds(lbs))")
	fmt.Scanln(&weightUnitfirst)
	if weightUnitfirst == 1 {
		fmt.Println("Choose a weight unit whoich your first unit will be converted to: (1 = pounds(lbs), )")
		fmt.Scanln(&weightUnitsecond)
		if weightUnitsecond == 1 {
			fmt.Println("Value of the first unit: ")
			fmt.Scanln(&weightValue)
			fmt.Println(weightValue*2.20462262, "lbs")
			fmt.Println("Press any key to quit program....")
			fmt.Scanln()
		}
	}
}

func Length() {

}

func main() {
	var typeUnit int
	fmt.Println("Choose your type of unit: (1 = weight, 2 = length)")
	fmt.Scanln(&typeUnit)
	switch {
	case typeUnit == 1:
		Weight()
	case typeUnit == 2:
		Length()
	}
}
