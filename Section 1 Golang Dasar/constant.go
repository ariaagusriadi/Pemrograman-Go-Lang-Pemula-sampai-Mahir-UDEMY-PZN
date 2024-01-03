package main

import "fmt"

func main() {
	const (
		firstName string = "aria"
		lastName         = "riadi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "agus"
	// firstName = "anjay"

}
