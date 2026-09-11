package main

/*
 * You can find external libraries from the [Go Package Repository](https://pkg.go.dev).
 * Find the package you need and add it to your `import` statement, using the name of the package as the import path.
 * To download the library for usage use the `go get` command.
 */
import (
	"fmt"
	"math"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Opt())

	var a any
	a = 5.4
	r := a.(float64)
	fmt.Println("the circles area is ", math.Pi*r*r)
}
