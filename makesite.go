package main

import (
	// "fmt"
	"bytes"
	"html/template"
	"io/ioutil"
)

func main() {
	// Reads file
	FileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	// define b as bytes, made to convert string into usable bytes to write to file
	var b bytes.Buffer
	// Set template
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(&b, string(FileContents))

	if err != nil {
		panic(err)
	}
	// Write to file
	err = ioutil.WriteFile("new-file.html", b.Bytes(), 777)

	if err != nil {
		panic(err) // This will fail if the file exists
		// Because it tries to write to a file that already exists
		// Solution: Store things inside of a folder that the user running the build has permissions to.
		// Link to solution: https://support.circleci.com/hc/en-us/articles/360003649774-Permission-Denied-When-Creating-Directory-or-Writing-a-File
	}

}