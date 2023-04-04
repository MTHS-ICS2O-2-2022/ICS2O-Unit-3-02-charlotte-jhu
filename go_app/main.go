// Created by: Charlotte Jhu
// Created on: April 2023
//
// This program calculates the volume of a square-based pyramid

package main

import "fmt"

func main() {
	// This function calculates the volume of a square-based pyramid
	// variables
	var length float64
	var width float64
	var height float64

	// input
	fmt.Println("This program calculates the volume of a square-based pyramid")
	fmt.Println()
	fmt.Print("Enter the length (mm): ")
	fmt.Scanln(&length)
	fmt.Print("Enter the width (mm): ")
	fmt.Scanln(&width)
	fmt.Print("Enter the height (mm): ")
	fmt.Scanln(&height)
	fmt.Println()

	// process
	volume := (length * width * height) / 3

	// output
	fmt.Println("The volume is ", volume, "mm³")
}
