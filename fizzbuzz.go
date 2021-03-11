package fizzbuzz

import "fmt"

func GenerateArray() []int {
	var range100 []int
	for i := 1; i <= 100; i++ {
		range100 = append(range100, i)
	}
	return range100
}

func Run(fbuzz []int) {
	range100 := fbuzz
	for i, value := range range100 {
		if value > 0 && value % 15 == 0 {
			PrintValues(i + 1, "FizzBuzz")
			} else if value > 0 && value % 5 == 0 {
				PrintValues(i + 1, "Buzz")
		} else if value > 0 &&  value % 3 == 0 {
			PrintValues(i + 1, "Fizz")
		} else {
			if i > 0 {
				fmt.Println(i)
			}
		}
	}
}

func PrintValues(index int, fizzBuzz string) {
	fmt.Printf("%v - %v\n", index, fizzBuzz)
}