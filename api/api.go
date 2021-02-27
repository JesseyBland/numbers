package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JesseyBland/numbers/app"
)

var (
	//Char is chacter comprising cnum
	Char string
)

//Server Starts Api Server with listen and handle functions.
func Server(port, char string) {
	Char = char
	http.HandleFunc("/", cnumAPI)
	http.ListenAndServe(":"+port, nil)
	fmt.Println("API Listening on port: " + port)
}

func cnumAPI(w http.ResponseWriter, r *http.Request) {

	pCnum, ok := r.URL.Query()["cnum"]
	pChar, ok2 := r.URL.Query()["char"]

	if !ok || len(pCnum[0]) < 1 {
		log.Println("Url Param 'Cnum' is missing")
		return
	}
	if !ok2 || len(pChar[0]) < 1 {
		log.Println("Url Param 'Char' is missing")
	}

	// Query()["pCnum"] will return an array of items,
	// we only want the single int 0 - 9999.
	realnumber, _ := strconv.Atoi(pCnum[0])
	res := app.Numbers(realnumber)
	if realnumber > 0 && realnumber < 10000 {

		if pChar != nil {
			Char = string(pChar[0])
		}
		res = app.Change(res, Char)
		fmt.Fprintf(w, res)
	} else {
		res = "Error out of range"
		fmt.Fprintf(w, res)
	}
	log.Println("Url Param 'number' is: " + string(pCnum[0]))
}
