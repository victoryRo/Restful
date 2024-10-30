package main

import (
	"flag"
	"fmt"
)

func main() {
	// flagMultiParams()
	// withInit()
	result := executeProgram()
	fmt.Println(result)

}

// func flagMultiParams() {
// 	var name = flag.String("name", "Stranger", "your wonderful name")
// 	var age = flag.Int("age", 100, "your graceful age")
//
// 	flag.Parse()
// 	log.Printf("Hellow %s (%d years), Welcome to the command line world", *name, *age)
// }

// -------------------------------------------

// var team string
// var points int
//
// func init() {
// 	flag.StringVar(&team, "team", "stranger", "your wonderful name")
// 	flag.IntVar(&points, "points", 0, "points of game...")
// }
// func withInit() {
// 	flag.Parse()
// 	log.Printf("Hi %s your points of the game are %d", team, points)
// }

// -------------------------------------------

var n1, n2 int
var operator string

func executeProgram() string {
	flag.StringVar(&operator, "operator", "any", "Describes the operation to be performed")
	flag.IntVar(&n1, "ope-1", 0, "Number parameter")
	flag.IntVar(&n2, "ope-2", 0, "Number parameter")

	flag.Parse()

	switch operator {
	case "sum":
		return fmt.Sprintf("With %s operator the result is: %.2f", operator, sum(n1, n2))
	case "mult":
		return fmt.Sprintf("With %s operator the result is: %.2f", operator, mult(n1, n2))
	case "div":
		return fmt.Sprintf("With %s operator the result is: %.2f", operator, div(n1, n2))
	default:
		return fmt.Sprintf("The operator %s Not Exists", operator)
	}
}

func sum(n1, n2 int) float32 {
	return float32(n1 + n2)
}

func mult(n1, n2 int) float32 {
	return float32(n1 * n2)
}

func div(n1, n2 int) float32 {
	return float32(n1 / n2)
}
