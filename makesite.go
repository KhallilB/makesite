package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {
	paths := []string{
		"todo.tmpl",
	}

	todos := []entry{
		entry{Name: "Task", Done: false},
	}
	user := ToDo{
		User: "Khallil",
		List: todos,
	}

	t := template.Must(template.New("todo.tmpl").ParseFiles(paths...))
	err := t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}

func readFile(src string) string {
	content, err := ioutil.ReadFile(src)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func writeFile(dst string, data []byte) error {
	return ioutil.WriteFile(dst, data, 0644)
}
