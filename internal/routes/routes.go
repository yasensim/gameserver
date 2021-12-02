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

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}

func Handlers() {
	http.HandleFunc("/welcome", welcome)
	http.Handle("/greet", http.HandlerFunc(greet))
	usersHandler := http.HandlerFunc(service.HandleUsers)
	userHandler := http.HandlerFunc(service.HandleUser)

	http.Handle("/users", CommonMiddleware(usersHandler))
	http.Handle("/user/", CommonMiddleware(userHandler))

}
