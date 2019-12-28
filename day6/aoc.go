package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type orbitMap map[string]string

func readInput(fileName string) (o []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		o = append(o, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func findOrbits(i []string) (orbits orbitMap) {
	orbits = make(orbitMap)
	for _, v := range i {
		objects := strings.Split(v, ")")
		orbits[objects[1]] = objects[0]
	}
	return
}

func (o orbitMap) countOrbits() (count int) {
	for object := range o {
		count += o.getIndirects(object)
	}
	return
}

// find Indirects till we can't find any
func (o orbitMap) getIndirects(object string) (count int) {
	for val, ok := o[object]; ok; val, ok = o[val] {
		count++
	}
	return
}

func (o orbitMap) getTravelDistance(from, to string) (distance int) {
	path := make(map[string]int)
	// Find the path of YOU and see what distance we travel per step
	for val, ok := o[from]; ok; val, ok = o[val] {
		path[val] = len(path)
	}
	// Find the path of SAN and see what distance we travel per step
	for val, ok := o[to]; ok; val, ok = o[val] {
		// See if it intersects with YOU and add it's distance to the total
		if d, found := path[val]; found {
			distance += d
			break
		}
		distance++
	}
	return
}

func main() {
	lines := readInput("input")
	orbits := findOrbits(lines)

	orbitCount := orbits.countOrbits()
	fmt.Printf("The awnser to question one is: %d\n", orbitCount)

	travelDistance := orbits.getTravelDistance("YOU", "SAN")
	fmt.Printf("The awnser to question two is: %d\n", travelDistance)
}
