package main

import (
	"log"

	"github.com/rhobro/utils.go/pkg/services/mongrel"
)

func init() {
	err := mongrel.Init("mongodb://root:groot@rhobro-tag-x6r67grr2pq9g-27017.githubpreview.dev/", "")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mongrel.C.
}
