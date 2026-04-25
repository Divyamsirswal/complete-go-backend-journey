package main

import "fmt"

/*
	Challenge 3:
		- Understand type conversion
			var x int = 10
			var y float64 = 20.5
			fmt.Println(x + int(y))
*/

func main() {
	var x int = 10
	var y float64 = 20.5

	/*
	  Here y  = 20.5 -> converts into -> 20 (for only calculation) because we have did int(y) but y = 20.5
	*/
	fmt.Println(x + int(y))

	// fmt.Println(y) -> try it
}
