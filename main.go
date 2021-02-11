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
	var (
		qq1 Quad
		qq2 Quad
		qq3 Quad
		qq4 Quad
	)

	switch {
	case n10 == "0":
		qq2 = q2two

	case n10 == "1":
		qq2 = q2one
	case n10 == "2":
		qq2 = q2two
	case n10 == "3":
		qq2 = q2three
	case n10 == "4":
		qq2 = q2four
	case n10 == "5":
		qq2 = q2five
	case n10 == "6":
		qq2 = q2six
	case n10 == "7":
		qq2 = q2seven
	case n10 == "8":
		qq2 = q2eight
	case n10 == "9":
		qq2 = q2nine

	}

	switch {
	case n1 == "0":
		//	q1=q1zero.l1,q1zero.l2,q1zero.l3,q1zero.l4
		qq1 = q1zero
	case n1 == "1":
		qq1 = q1one
	//	q1=q1one.l1,q1onel2,q1one.l3,q1one.l4
	case n1 == "2":
		qq1 = q1two
	//	q1=q1two.l1,q1two.l2,q1two.l3,q1two.l4
	case n1 == "3":
		qq1 = q1three
	//	q1=q1three.l1,q1three.l2,q1three.l3,q1three.l4
	case n1 == "4":
		qq1 = q1four
	//	q1=q1four.l1,q1four.l2,q1four.l3,q1four.l4
	case n1 == "5":
		qq1 = q1five
	//	q1=q1five.l1,q1five.l2,q1five.l3,q1five.l4
	case n1 == "6":
		qq1 = q1six
	//	q1=q1six.l1,q1six.l2,q1six.l3,q1six.l4
	case n1 == "7":
		qq1 = q1seven
	//	q1=q1seven.l1,q1seven.l2,q1seven.l3,q1seven.l4
	case n1 == "8":
		qq1 = q1eight
	//	q1=q1eight.l1,q1eight.l2,q1eight.l3,q1eight.l4
	case n1 == "9":
		qq1 = q1nine
		//	q1=q1nine.l1,q1nine.l2,q1nine.l3,q1nine.l4
	}
	switch {
	case n1000 == "0":
		qq4 = q4zero
	case n1000 == "1":
		qq4 = q4one
	case n1000 == "2":
		qq4 = q4two
	case n1000 == "3":
		qq4 = q4three
	case n1000 == "4":
		qq4 = q4four
	case n1000 == "5":
		qq4 = q4five
	case n1000 == "6":
		qq4 = q4six
	case n1000 == "7":
		qq4 = q4seven
	case n1000 == "8":
		qq4 = q4eight
	case n1000 == "9":
		qq4 = q4nine

	}

	switch {
	case n100 == "0":
		qq3 = q3two
	case n100 == "1":
		qq3 = q3one
	case n100 == "2":
		qq3 = q3two
	case n100 == "3":
		qq3 = q3three
	case n100 == "4":
		qq3 = q3four
	case n10 == "5":
		qq3 = q3five
	case n100 == "6":
		qq3 = q3six
	case n100 == "7":
		qq3 = q3seven
	case n100 == "8":
		qq3 = q3eight
	case n100 == "9":
		qq3 = q3nine

	}

	// q2q2q2 hl q1q1q1
	// q2q2q2 hl q1q1q1
	// q2q2q2 hl q1q1q1
	//  hxhxhxhxhxhxhx
	//  hxhxhxhxhxhxhx
	// q3q3q3 hl q4q4q4
	// q3q3q3 hl q4q4q4
	// q3q3q3 hl q4q4q4

	fmt.Println("*****************NUMBER****************")
	fmt.Println(qq2.l1 + hl + qq1.l1)
	fmt.Println(qq2.l2 + hl + qq1.l2)
	fmt.Println(qq2.l3 + hl + qq1.l3)
	fmt.Println(qq2.l4 + hl + qq1.l4)

	fmt.Println(hx)
	fmt.Println(hx)

	fmt.Println(qq3.l1 + hl + qq4.l1)
	fmt.Println(qq3.l2 + hl + qq4.l2)
	fmt.Println(qq3.l3 + hl + qq4.l3)
	fmt.Println(qq3.l4 + hl + qq4.l4)

	fmt.Println("**************************************")

	fmt.Println("\n\n*****", good, "t=", t, "number=", number, "n1000=", n1000, "n100=", n100, "n10=", n10, "n1=", n1)
}
