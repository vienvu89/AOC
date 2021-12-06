package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func partA(a []string) {
	boards :=
}

func partB(a []string) {

}

type item struct {
	number int
	marked bool
}

func main() {
	data, err := ioutil.ReadFile("test.txt")
	//data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading eroor", err)
		return
	}

	a := strings.Split(string((data)), "\n\n")
	fmt.Println(len(a))
}
