package app

import (
	"fmt"
	"strconv"
)

//SetPort sets the port for parameter port
func SetPort(port string) string {

	portInt, _ := strconv.Atoi(port)
	if portInt >= 0 && portInt <= 9999 {
		return port
	}
	fmt.Println("PORT OUT OF RANGE")
	return "9000"

}
