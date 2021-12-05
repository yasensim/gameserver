package main

import (
	"log"
	"net/http"

	"github.com/yasensim/gameserver/internal/routes"
)

func main() {
	r := routes.Handlers()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
