package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func partA(b []int) {
	count := 0
	for i := 1; i < len(b); i++ {
		if b[i] > b[i - 1] {
			count += 1
		}
	}
	fmt.Println(count)
}

func partB(b []int) {
	count := 0

	for i := 0; i < len(b) - 3; i++ {
		sum1 := b[i] + b[i + 1] + b[i + 2]
		sum2 := b[i + 1] + b[i + 2] + b[i + 3]
		if sum2 > sum1 {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	s := strings.Split(string(data), "\n")
	b := make([]int, len(s))
	for i, v := range s {
		//fmt.Println(v)
		b[i], _ = strconv.Atoi(v)
	}
	//partA(b)
	partB(b)
}
