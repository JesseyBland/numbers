package app

import (
	"strconv"
)

//Local produces a cnum 0-9999 through the console
func Local(number, char string) string {
	realnumber, _ := strconv.Atoi(number)
	if realnumber > 0 && realnumber < 10000 {
		res := Numbers(realnumber)
		res = Change(res, char)
		return res
	}
	res := "Error out of range"
	return res

}
