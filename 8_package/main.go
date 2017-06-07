package main

import (
	"fmt"

	"github.com/hfogelberg/tempconv"
)

func main() {
	f := tempconv.Fahrenheit(100)
	c := tempconv.FToC(f)
	fmt.Println(c.String())

}
