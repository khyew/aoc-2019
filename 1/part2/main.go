package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// fuelRequired returns the amount of fuel needed for a module of mass m
func fuelRequired(m int) int {
	return m/3 - 2
}

func main() {
	f, err := os.Open("input")

	if err != nil {
		fmt.Printf("Couldn't open input file: %v\n", err)
		return
	}

	totalFuel := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		mass, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Couldn't parse line (%v) into integer: %v\n", line, err)
		}

		fuelMass := fuelRequired(mass)

		for fuelMass > 0 {
			totalFuel += fuelMass
			fuelMass = fuelRequired(fuelMass)
		}
	}

	fmt.Printf("Total fuel required: %v\n", totalFuel)
}
