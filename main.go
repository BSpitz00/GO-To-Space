package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fuel := 1000000
	fmt.Println("What planet would you like to visit?")
	fmt.Println("Choices: Venus, Mercury, Mars")

	planetChoice := "" //Enter choice here

	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
	cost(fuel)
	timothySmells()

}

func fuelGauge(fuel int) {
	fmt.Println("You have", fuel, "gallons of fuel left!")
}

func calculateFuel(planet string) int {
	switch planet {
	case "Venus":
		return 300000
	case "Mercury":
		return 500000
	case "Mars":
		return 700000
	default:
		return 0
	}
}

func greetPlanet(planet string) {
	fmt.Printf("Welcome to planet %v", planet)
}

func cantFly() {
	fmt.Print("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	fuelRemaining := fuel
	fuelCost := calculateFuel(planet)

	if fuelRemaining < fuelCost {
		cantFly()
		return -1
	}

	greetPlanet(planet)
	fuelRemaining -= fuelCost

	return fuelRemaining - fuelCost
}

func cost(fuel int) {
	cost := fuel * 4
	fmt.Printf("The cost of fuel for this trip is $%v... Thanks Mr. President", cost)
}

func timothySmells() {
	rand.Seed(time.Now().UnixNano())
	aliens := rand.Intn(1000)
	fmt.Printf("%v inhabitants on planet think Timothy smells", aliens)
}
