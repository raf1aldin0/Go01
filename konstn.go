package main

import "fmt"

func main() {

	const lastName = "wick"
	const firstName string = "john"
	fmt.Print("nice to meet you ", lastName, "!\n", "halo ", firstName, "!\n")

	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

}
