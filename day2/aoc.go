package main

import "fmt"

/////
// DAY 2
/////
func intCode(a codeArray) int {
	i := 0
	for {
		opcode := a[i]
		if opcode == 1 {
			value1 := a[a[i+1]]
			value2 := a[a[i+2]]
			a[a[i+3]] = value1 + value2
		} else if opcode == 2 {
			value1 := a[a[i+1]]
			value2 := a[a[i+2]]
			a[a[i+3]] = value1 * value2
		} else if opcode == 99 {
			return a[0]
		} else {
			panic(fmt.Sprint(a[i], a[1], a[2]))
		}
		i += 4
	}
}

type codeArray []int

func (a codeArray) restoreGravity(n int, v int) codeArray {
	a[1] = n
	a[2] = v
	return a
}

func main() {
	baseArray := codeArray{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 2, 9, 19, 23, 1, 9, 23, 27, 2, 27, 9, 31, 1, 31, 5, 35, 2, 35, 9, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 47, 6, 51, 2, 51, 10, 55, 1, 9, 55, 59, 2, 6, 59, 63, 1, 63, 6, 67, 1, 67, 10, 71, 1, 71, 10, 75, 2, 9, 75, 79, 1, 5, 79, 83, 2, 9, 83, 87, 1, 87, 9, 91, 2, 91, 13, 95, 1, 95, 9, 99, 1, 99, 6, 103, 2, 103, 6, 107, 1, 107, 5, 111, 1, 13, 111, 115, 2, 115, 6, 119, 1, 119, 5, 123, 1, 2, 123, 127, 1, 6, 127, 0, 99, 2, 14, 0, 0}

	a := make(codeArray, len(baseArray))
	b := make(codeArray, len(baseArray))
	copy(a, baseArray)

	a.restoreGravity(12, 2)
	fmt.Printf("The awnser to question one is: %d\n", intCode(a))

	goal := 19690720
	var noun, verb int
	for i := 0; i < 1000000; i++ {
		copy(b, baseArray)

		noun = int(i / 100)
		verb = i - (100 * noun)

		b.restoreGravity(noun, verb)

		if intCode(b) == goal {
			break
		}
	}

	fmt.Printf("The awnser to question two is: %d\n", 100*noun+verb)

}
