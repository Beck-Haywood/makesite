package main

import (
	"fmt"
	"flag"
	"os"
	"html/template"
	"io/ioutil"
	"strings"

	"context"
	"cloud.google.com/go/translate"
	_ "cloud.google.com/go/translate/apiv3"
	"golang.org/x/text/language"
)

type Post struct {
	Text string
}

func main() {
	// Get flag value
	filePath := flag.String("postPath", "first-post.txt", "Name of file you want to read from.")
	outputPath := flag.String("outputPath", "new-file.html", "Name of file you want to output to.")
	dirName := flag.String("dir", ".", "This is the directory.")
	lang := flag.String("lang", "es", "This is the language you want to translate, inputting google's language abbreviations.")

	flag.Parse()
	fmt.Println("Language:", lang)

	files, err := ioutil.ReadDir(*dirName)
	if err != nil {
		panic(err)
	}
	tail := "txt"
	for _, file := range files {
		for i := range file.Name() {
			if file.Name()[i] == '.' {
				s := strings.Split(file.Name(), ".")[1]
				if s == tail {
					// Is txt file
					fmt.Println(file.Name())
					fileContents, err := ioutil.ReadFile(file.Name())
					if err != nil {
						panic(err)
					}

					contents, error := translateText(*lang, string(fileContents))
					if error != nil {
						panic(error)
					}
					bytesToWrite := []byte(contents)

					err1 := ioutil.WriteFile(file.Name(), bytesToWrite, 0644)
					if err1 != nil {
						panic(err1)
					}
					writeHTMLGivenFile("template.tmpl", file.Name())
				}
			}
		}
	}

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

func writeHTMLGivenFile(templateName string, fileName string) {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	post := Post{Text: string(fileContents)}
	t := template.Must(template.New("template.tmpl").ParseFiles(templateName))
	ext := ".html"
	filter := strings.Split(fileName, ".")[0] + ext
	f, err := os.Create(filter)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, post)
	if err != nil {
		panic(err)
	}
}

func translateText(targetLanguage, text string) (string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	// fmt.Println(resp[0].Text, nil)
	return resp[0].Text, nil
}