package main

import (
	"fmt"
	"strings"
)

/////
// DAY 4
/////

func intToArray(i int) (a [6]int) {
	a[0] = int(i / 100000)
	i -= a[0] * 100000
	a[1] = int(i / 10000)
	i -= a[1] * 10000
	a[2] = int(i / 1000)
	i -= a[2] * 1000
	a[3] = int(i / 100)
	i -= a[3] * 100
	a[4] = int(i / 10)
	i -= a[4] * 10
	a[5] = i

	return
}

func adjacentDigits(i int) (foundAdjacent bool) {
	intArray := intToArray(i)
	for k, v := range intArray {
		if k == 0 {
			continue
		}
		if v == intArray[k-1] {
			foundAdjacent = true
			break
		}
	}
	return
}

func adjacentDigitsGroups(i int) (foundAdjacent bool) {
	intArray := intToArray(i)
	for k, v := range intArray {
		if k == 0 {
			continue
		}
		if v == intArray[k-1] {
			if strings.Count(fmt.Sprint(i), fmt.Sprint(v)) == 2 {
				foundAdjacent = true
				break
			}
		}
	}
	return
}

func decreased(i int) (decreased bool) {
	intArray := intToArray(i)
	for k, v := range intArray {
		if k == 0 {
			continue
		}
		if v < intArray[k-1] {
			decreased = true
		}
	}
	return
}

func findPasswordRange1(l int, u int) (pos int) {
	for i := l; i < u; i++ {
		if adjacentDigits(i) && !decreased(i) {
			pos++
		}
	}
	return
}

func findPasswordRange2(l int, u int) (pos int) {
	for i := l; i < u; i++ {
		if adjacentDigitsGroups(i) && !decreased(i) {
			pos++
		}
	}
	return
}

func main() {

	lowerLimit := 307237
	upperLimit := 769058

	awnser1 := findPasswordRange1(lowerLimit, upperLimit)
	fmt.Printf("The awnser to question one is: %d\n", awnser1)

	awnser2 := findPasswordRange2(lowerLimit, upperLimit)
	fmt.Printf("The awnser to question one is: %d\n", awnser2)
}
