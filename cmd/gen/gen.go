package main

import (
	"errors"
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
	BC        string
}

func main() {
	var d data
	var typeGen string

	flag.StringVar(&typeGen, "type", "", "Query/Command")
	flag.StringVar(&d.Name, "name", "", "The name in camelcase [BC][SERVICE]")
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
	if len(splitted) < 2 {
		fmt.Println("Name format invalid need to be [BC][Service]")
	}

	d.Name = splitted[1]
	d.NameLower = strings.ToLower(splitted[1])
	d.BC = strings.ToLower(splitted[0])

	bc := d.BC
	service := d.NameLower

	selectedType := strings.ToLower(typeGen)
	d.NameLower = strings.ToLower(d.Name)

	files := []string{"handler.tmpl", "http_handler.tmpl", "http_handler_test.tmpl", "service.tmpl", "service_test.tmpl"}
	filesToCreate := make(map[string]string)
	filesToCreate["http_handler.tmpl"] = fmt.Sprintf("%s/%s.go", routeHandler(service), service)
	filesToCreate["http_handler_test.tmpl"] = fmt.Sprintf("%s/%s_test.go", routeHandler(service), service)
	filesToCreate["service.tmpl"] = fmt.Sprintf("%s/service.go", routeService(bc)(service))
	filesToCreate["service_test.tmpl"] = fmt.Sprintf("%s/service_test.go", routeService(bc)(service))
	filesToCreate["handler.tmpl"] = fmt.Sprintf("%s/%s.go", routeService(bc)(service), selectedType)

	createFolder(bc, routeBoundedContext)
	createFolder(service, routeService(bc))
	createFolder(service, routeHandler)

	for _, file := range files {
		f, err := os.Create(filesToCreate[file])
		defer func() {
			err := f.Close()
			if err != nil {
				fmt.Println("error closing file")
			}
		}()
		if err != nil {
			panic(fmt.Sprintf("error: %w", err))
		}
		t := template.Must(template.ParseFiles(fmt.Sprintf("cmd/gen/templates/%s/%s", selectedType, file)))
		t.Execute(f, d)
	}
}

func createFolder(folder string, f func(string) string) {
	dirName := f(folder)
	_, err := os.Stat(dirName)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(f(folder), 0755)
		if err != nil {
			panic(fmt.Sprintf("can't create the directory error: %s", err.Error()))
		}
	}
}

func routeHandler(service string) string {
	return fmt.Sprintf("%s/platform/server/handler/%s", routeInternal(), service)
}

func routeService(bc string) func(string) string {
	return func(service string) string {
		return fmt.Sprintf("%s/%s", routeBoundedContext(bc), service)
	}
}

func routeBoundedContext(bc string) string {
	return fmt.Sprintf("%s/%s", routeInternal(), bc)
}

func routeInternal() string {
	return "internal"
}
