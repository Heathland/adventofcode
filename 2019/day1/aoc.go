package main

import (
	"fmt"
	"math"
)

///
// DAY 1
///
func calcFuelRecursion(i float64) (fuel float64) {
	for i >= 0 {
		i = calcFuel(i)
		if i >= 0 {
			fuel += i
		}
	}
	return
}

func calcFuel(i float64) float64 {
	return math.Floor(float64(i)/3) - 2
}

func main() {
	a := []float64{58444, 100562, 133484, 67910, 58372, 104607, 108786, 137410, 62910, 76115, 64142, 59324, 54327, 92864, 94120, 63931, 128696, 111758, 65698, 54930, 116136, 127111, 133914, 52992, 90364, 107637, 62118, 147901, 62347, 53614, 140690, 115587, 66148, 95729, 148847, 84269, 71569, 85026, 130871, 102470, 53328, 63308, 104085, 57744, 123008, 120983, 94968, 69402, 83830, 137069, 121062, 71267, 103035, 97604, 129153, 65595, 148655, 124573, 139257, 59722, 101050, 139557, 74362, 50024, 101750, 83209, 117840, 139442, 127810, 113438, 94731, 125471, 96653, 88522, 125573, 74456, 89839, 84458, 128451, 68608, 92504, 103549, 117980, 126850, 144675, 59752, 60986, 125867, 89982, 108717, 134815, 89209, 143434, 61123, 103162, 139529, 122228, 137866, 78676, 80779}

	var sumFuelOne, sumFuelTwo float64
	for _, v := range a {
		sumFuelOne += calcFuel(v)
	}
	fmt.Printf("The awnser to question one is: %.f\n", sumFuelOne)

	for _, v := range a {
		sumFuelTwo += calcFuelRecursion(v)
	}
	fmt.Printf("The awnser to question two is: %.f\n", sumFuelTwo)
}
