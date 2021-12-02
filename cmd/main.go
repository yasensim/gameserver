package main

import (
	"log"
	"net/http"

	"github.com/yasensim/gameserver/internal/routes"
)

func main() {
	routes.Handlers()

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
