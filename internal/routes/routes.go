package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yasensim/gameserver/internal/users/auth"
	usersService "github.com/yasensim/gameserver/internal/users/service"
)

//func printHeaders(r *http.Request) {
//	fmt.Printf("Request at %v\n", time.Now())
//	for k, v := range r.Header {
//		fmt.Printf("%v: %v\n", k, v)
//	}
//}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Req from middleware: %s \n", r.Header.Get("test"))
		if len(r.Header.Get("test")) > 0 {
			w.Header().Add("test", r.Header.Get("test"))
		}
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}

func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	us := usersService.Get()
	jv := auth.GetAuthenticator()

	r.HandleFunc("/register", us.CreateUser).Methods("POST")
	r.HandleFunc("/login", us.Login).Methods("POST")

	s := r.PathPrefix("/auth").Subrouter()
	s.Use(jv.JwtVerify)

	s.HandleFunc("/user", us.FetchUsers).Methods("GET")
	s.HandleFunc("/user/{id}", us.GetUser).Methods("GET")
	s.HandleFunc("/user/{id}", us.UpdateUser).Methods("PUT")
	s.HandleFunc("/user/{id}", us.DeleteUser).Methods("DELETE")
	return r

}
