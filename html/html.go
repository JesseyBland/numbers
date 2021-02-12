package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func cnumber(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}
func headers(w http.ResponseWriter, req *http.Request) {
	mybody := req.Body
	buffer := make([]byte, 1024)
	myint, _ := mybody.Read(buffer)
	snum, _ := strconv.Atoi(string(buffer[:myint]))
	resp := numbers(snum)
	fmt.Fprintf(w, resp)
}

// func headers(w http.ResponseWriter, req *http.Request) {
// 	// mybody := req.Body
// 	// buffer := make([]byte, 1024)
// 	// mybody.Read(buffer)
// 	// fmt.fPrintf(w, string(buffer))

// 	// fmt.Fprintf(w, "POST request successful\n")
// 	// number := req.FormValue("number")

// 	// fmt.Fprintf(w, "Number = %s\n", number)

// }

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/cnumber", cnumber)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":9000", nil)
}
