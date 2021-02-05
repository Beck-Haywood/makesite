package main

import (
	// "fmt"
	"flag"
	"os"
	// "bytes"
	"html/template"
	"io/ioutil"

)

type Post struct {
	Text string
}

func main() {
	// Get flag value
	filePath := flag.String("postPath", "first-post.txt", "Name of file you want to read from.")
	outputPath := flag.String("outputPath", "new-file.html", "Name of file you want to output to.")
	flag.Parse()

	// Read file using flag 
	fileContents, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}
	
	post := Post{string(fileContents)}

	// Init template
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// Create output file if it does not exist
	file, err := os.Create(*outputPath)
	// This will fail if the file ("new-file.html") exists
		// Because it tries to write to a file that already exists
		// Solution: Store things inside of a folder that the user running the build has permissions to.
		// Link to solution: https://support.circleci.com/hc/en-us/articles/360003649774-Permission-Denied-When-Creating-Directory-or-Writing-a-File
	if err != nil{
		panic(err)
	}

	errExecute := t.Execute(file, post)
	if errExecute != nil {
		panic(errExecute)
	}

	file.Close()

}