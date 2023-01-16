package main

import (
    "fmt"
    "math/rand"
    "time"
)

//Main is always the first method
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
    //You don't need a variable here, you are not modifying it after the switch, so return in the switch
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
    //A default will ALWAYS run if nothing else triggers, so no need for an external return statment
}

func greetPlanet(planet string) {
    //We use formatting (https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) instead of the java way of inserting strings mid string
    fmt.Printf("Welcome to planet %v", planet)
}

func cantFly() {
    fmt.Print("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
    //We do not need to create variables and set them to something else right after
    //You also do not need to strictly cast a variable, GoLang uses inferred variable types
    //with the := operator
    fuelRemaining := fuel
    fuelCost := calculateFuel(planet)

    if fuelRemaining < fuelCost {
        cantFly()
        return -1
    }

    //Just cleaned up to use return breaks, rather than else ifs, makes the code easier to read and segment
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