package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fuel := 1000000

	for fuel > 0 {
		fmt.Println("\nWhat planet in our solar system would you like to visit?")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		planetChoice := scanner.Text()

		fuel = flyToPlanet(planetChoice, fuel)
		fuelGauge(fuel)
		cost(fuel)
		youSmell()
	}
}

func fuelGauge(fuel int) {
	fmt.Println("you have", fuel, "gallons of fuel left!")
}

func calculateFuel(planet string) int {
	switch planet {
	case "The Sun":
		return 93000
	case "Mercury":
		return 58000
	case "Venus":
		return 26000
	case "The Moon":
		return 1000
	case "Mars":
		return 49000
	case "Jupiter":
		return 391000
	case "Saturn":
		return 796000
	case "Uranus":
		return 1697000
	case "Neptune":
		return 2787000
	case "Pluto":
		return 3577000
	default:
		return 0
	}
}

func greetPlanet(planet string) {
	fmt.Printf("Welcome to planet %v, ", planet)
}

func cantFly() {
	fmt.Print("Too far! We do not have the available fuel to fly there, you're stuck!")
}

func flyToPlanet(planet string, fuel int) int {
	fuelRemaining := fuel
	fuelCost := calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func cost(fuel int) {
	cost := fuel * 4
	fmt.Printf("The cost of fuel for this trip is $%v.", cost)
}

func youSmell() {
	rand.Seed(time.Now().UnixNano())
	aliens := rand.Intn(1000)
	fmt.Printf(" %v Inhabitants on planet think you smell funny.", aliens)
}
