package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/yasensim/gameserver/internal/users/service"
)

func printHeaders(r *http.Request) {
	fmt.Printf("Request at %v\n", time.Now())
	for k, v := range r.Header {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	printHeaders(r)
	w.Write([]byte("welcome to go!!!"))
}

func greet(w http.ResponseWriter, r *http.Request) {
	printHeaders(r)
	w.Write([]byte("Hello from go"))
}

func Handlers() {
	http.HandleFunc("/welcome", welcome)
	http.Handle("/greet", http.HandlerFunc(greet))
	usersHandler := http.HandlerFunc(service.HandleUsers)
	userHandler := http.HandlerFunc(service.HandleUser)
	http.Handle("/users", usersHandler)
	http.Handle("/user/", userHandler)

}
