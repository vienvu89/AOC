package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func partA(a []string) {
	horizontal := 0
	depth := 0
	for _, s := range (a) {
		b := strings.Split(s, " ")
		h, err := strconv.Atoi(b[1])
		if err != nil {
			fmt.Println("Error:", err)
		}
		switch b[0] {
		case "forward":
			horizontal += h
		case "up":
			depth -= h
		case "down":
			depth  += h
		default:
			fmt.Println("Error:")
		}
	}
	fmt.Println(horizontal * depth)
}

func partB(a []string) {
	horizontal := 0
	depth := 0
	aim := 0
	for _, s := range (a) {
		b := strings.Split(s, " ")
		h, err := strconv.Atoi(b[1])
		if err != nil {
			fmt.Println("Error:", err)
		}
		switch b[0] {
		case "forward":
			horizontal += h
			depth += h * aim
		case "up":
			aim -= h
		case "down":
			aim += h
		default:
			fmt.Println("Error:")
		}
	}
	fmt.Println(horizontal * depth)
}

func main() {
	//data, err := ioutil.ReadFile("test.txt")
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading eroor", err)
		return
	}
	a := strings.Split(string((data)), "\n")
	partB(a)
}
