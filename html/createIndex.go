package html

import (
	"io/ioutil"
	"log"
)

//creates a temp index.html from var index
func createindex() string {
	file, err := ioutil.TempFile(".", "index.*.html")
	file.WriteString(index)
	if err != nil {
		log.Fatal(err)
	}
	file.Sync()
	return file.Name()
}
