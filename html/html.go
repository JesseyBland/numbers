package html

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/JesseyBland/numbers/app"
)

var (
	//Char sets the character stated in the -char flag
	Char string
)

//serves the index html page
func static(w http.ResponseWriter, r *http.Request) {
	file := (createindex())
	page, err := template.ParseFiles(file)
	if err != nil {
		fmt.Println(err)
	}

	page.Execute(w, nil)
	defer os.Remove(file)
}

//function returns the cnum when given the 0-9999 string
func cnumber(w http.ResponseWriter, req *http.Request) {

	mybody := req.Body
	buffer := make([]byte, 1024)
	myint, _ := mybody.Read(buffer)
	snum, _ := strconv.Atoi(string(buffer[:myint]))
	resp := app.Numbers(snum)
	resp = app.Change(resp, Char)
	fmt.Fprintf(w, resp)

}

//Server is the HTML SERVER
func Server(port, char string) {
	Char = char
	http.HandleFunc("/", static)
	http.HandleFunc("/cnumber", cnumber)
	fmt.Println("HTML Listening on port: " + port)
	http.ListenAndServe(":"+port, nil)

}
