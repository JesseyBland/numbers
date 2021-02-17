package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func cnumber(w http.ResponseWriter, req *http.Request) {
	mybody := req.Body
	buffer := make([]byte, 1024)
	myint, _ := mybody.Read(buffer)
	snum, _ := strconv.Atoi(string(buffer[:myint]))
	resp := numbers(snum)
	fmt.Fprintf(w, resp)
}

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/cnumber", cnumber)
	fmt.Println("Listening on port: 9000")
	http.ListenAndServe(":9000", nil)

}
