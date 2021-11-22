package main

import (
	"log"

	"github.com/arkiant/ddd-golang-framework/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
