package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", cnumApi)
	http.ListenAndServe(":9001", nil)
}

func cnumApi(w http.ResponseWriter, r *http.Request) {

	pCnum, ok := r.URL.Query()["cnum"]

	if !ok || len(pCnum[0]) < 1 {
		log.Println("Url Param 'Cnum' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	realnumber, _ := strconv.Atoi(pCnum[0])
	if realnumber > 0 && realnumber < 9999 {
		res := numbers(realnumber)
		fmt.Fprintf(w, res)
	} else {
		res := "Error out of Range"
		fmt.Fprintf(w, res)
	}
	log.Println("Url Param 'number' is: " + string(pCnum[0]))
}
