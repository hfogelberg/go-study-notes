/**
Flags are a way to specify options for command-line programs.
If a flag is omitted it automatcally takes the default value.
*/

package main

import (
	"flag"
	"fmt"
	"strings"
)

// Arguments: name of flag, default value, messag to user
// if and invalid value is passed
var n = flag.Bool("n", false, "omit trailing spaces")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
