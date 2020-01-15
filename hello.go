package main

import  (
	"fmt"
)


func main() {
	// declaration directly typed
	//var a int = 5

	// shorthand syntax
	b := 5

	fmt.Println(b)

	if b > 0 {
		fmt.Println("Higher than")
	}

	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	//simple array
	var a[7] int
	a[2] = 9
	fmt.Println(a)

	//slice 
	c := []int {1, 2, 3, 4, 5}
	c = append(c, 6)
	fmt.Println(c)

	vertices := make(map[string] int)

	vertices["triangule"] = 3
	vertices["square"] = 4

	fmt.Println(vertices)
	fmt.Println(vertices["square"])
}
