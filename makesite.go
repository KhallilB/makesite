package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type data struct {
	Text string
}

func getFileName() (string, error) {
	var file string

	flag.StringVar(&file, "file", "", "File Name")
	flag.Parse()

	return file, nil
}

func readFile(text string) (string, error) {
	content, err := ioutil.ReadFile(text)

	if err != nil {
		panic(err)
	}

	return string(content), nil
}

func writeFile(dst string, data []byte) error {
	return ioutil.WriteFile(dst, data, 0644)
}

func createTemplate(text string, writer io.Writer) (string, error) {
	content := data{Text: text}

	tmpl, err := template.ParseFiles("template.html")

	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer

	if err := tmpl.Execute(&buffer, content); err != nil {
		return "", fmt.Errorf("error rendering data into template")
	}

	return buffer.String(), nil
}

func writeToTemplate(text string, writer io.Writer) error {
	content := data{Text: text}

	tmpl, err := template.ParseFiles("template.html")

	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, content); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	file, err := getFileName()

	data, err := readFile(file)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = writeToTemplate(data, w)

		if err != nil {
			log.Fatal("Could could not render data into html.")
		}
	})

	http.ListenAndServe(":2525", nil)
}
