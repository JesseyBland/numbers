package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	number = 10000
	var n1000 string
	var n100 string
	var n10 string
	var n1 string
	bad := "Error"
	good := "Success"
	resp := "Please enter a number 0-9999"

	for number < 0 || number > 9999 {
		fmt.Println(resp)
		fmt.Scan(&number)
		if number < 0 || number > 9999 {
			fmt.Println(bad)
		}
	}
	t := strconv.Itoa(number)

	if number < 1000 {
		n1000 = "0"
		n100 = string(t[0])
		n10 = string(t[1])
		n1 = string(t[2])
	} else if number < 100 {
		n1000 = "0"
		n100 = "0"
		n10 = string(t[0])
		n1 = string(t[1])
	} else if number < 10 {
		n1000 = "0"
		n100 = "0"
		n10 = "0"
		n1 = string(t[0])

	} else {
		n1000 = string(t[0])
		n100 = string(t[1])
		n10 = string(t[2])
		n1 = string(t[3])

	}

	fmt.Println(good, "t=", t, number, n1000, n100, n10, n1)
}
