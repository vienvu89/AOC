package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func partA(a []string) {
	track := make([]int, len(a[0]))
	for _,s := range(a) {
		b := strings.Split(s, "")
		for i, n := range (b) {
			if  n == "1" {
				track[i] += 1
			}
		}
	}
	gamma := ""
	epsilon := ""
	for _, n := range (track) {
		if n > len(a) / 2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaInt := int64(0)
	epsilonInt := int64(0)
	fmt.Println(gamma)
	if i, err := strconv.ParseInt(gamma, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		gammaInt = i
	}

	if i, err := strconv.ParseInt(epsilon, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		epsilonInt = i
	}

	fmt.Println(gammaInt * epsilonInt)
}

func partB(a []string) {


	copy := a
	i := 0
	for len(copy) > 1 {
		oxy := make([]string, 0, len(a))
		co2 :=  make([]string, 0, len(a))

 		for _, s := range(copy) {
			b := strings.Split(s, "")
			c := b[i]
			if c == "1" {
				oxy = append(oxy, s)
			} else {
				co2 = append(co2, s)
			}
		}
		if len(oxy)  >= len(co2) {
			copy = oxy
		} else {
			copy = co2
		}
		i += 1
	}
	fmt.Println(copy)

	oxyInt := int64(0)
	if i, err := strconv.ParseInt(copy[0], 2, 64); err != nil {
		fmt.Println(err)
	} else {
		oxyInt = i
	}

	copy = a
	i = 0
	for len(copy) > 1 {
		oxy := make([]string, 0, len(a))
		co2 :=  make([]string, 0, len(a))

		for _, s := range(copy) {
			b := strings.Split(s, "")
			c := b[i]
			if c == "1" {
				oxy = append(oxy, s)
			} else {
				co2 = append(co2, s)
			}
		}
		if len(oxy)  >= len(co2) {
			copy = co2
		} else  {
			copy = oxy
		}
		i += 1
	}

	co2Int := int64(0)
	if i, err := strconv.ParseInt(copy[0], 2, 64); err != nil {
		fmt.Println(err)
	} else {
		co2Int = i
	}

	fmt.Println(oxyInt * co2Int)
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
