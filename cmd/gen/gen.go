package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/fatih/camelcase"
)

type data struct {
	Name      string
	NameLower string
}

func main() {
	var d data
	var typeGen string

	flag.StringVar(&typeGen, "type", "", "Query/Command")
	flag.StringVar(&d.Name, "name", "", "The name")
	flag.Parse()

	if typeGen == "" {
		fmt.Println("Type need to be set")
		return
	}

	if d.Name == "" {
		fmt.Println("Name need to be set")
		return
	}

	splitted := camelcase.Split(d.Name)
	splittedName := strings.ToLower(strings.Join(splitted, "_"))

	selectedType := strings.ToLower(typeGen)
	d.NameLower = strings.ToLower(d.Name)

	files := []string{"handler.tmpl", "http_handler.tmpl", "service.tmpl"}
	root := "internal"
	filesToCreate := make(map[string]string)
	filesToCreate["handler.tmpl"] = fmt.Sprintf("%s/%s/%s.go", root, splittedName, selectedType)
	filesToCreate["http_handler.tmpl"] = fmt.Sprintf("%s/platform/server/handler/%s.go", root, splittedName)
	filesToCreate["service.tmpl"] = fmt.Sprintf("%s/%s/service.go", root, splittedName)

	// Create service folder
	err := os.Mkdir(fmt.Sprintf("%s/%s", root, splittedName), 0755)
	if err != nil {
		panic(fmt.Sprintf("can't create the directory error: %s", err.Error()))
	}

	for _, file := range files {
		f, err := os.Create(filesToCreate[file])
		defer func() {
			err := f.Close()
			if err != nil {
				fmt.Println("error closing file")
			}
		}()
		if err != nil {
			panic("error")
		}
		t := template.Must(template.ParseFiles(fmt.Sprintf("cmd/gen/templates/%s/%s", selectedType, file)))
		t.Execute(f, d)
	}
}
