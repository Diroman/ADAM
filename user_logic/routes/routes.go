package routes

import (
	"auth/controllers"
	"auth/utils/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.HandleFunc("/", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	// User operation
	s.HandleFunc("/user", controllers.FetchUsers).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	// favorite operation
	s.HandleFunc("/favorite", controllers.CreateFavorite).Methods("POST")
	s.HandleFunc("/favorite/{id}", controllers.GetFavorite).Methods("GET")
	s.HandleFunc("/favorite/user/{id}", controllers.GetAllFavorite).Methods("GET")
	s.HandleFunc("/favorite/{id}", controllers.DeleteFavorite).Methods("DELETE")
	// history operation
	s.HandleFunc("/history", controllers.CreateHistory).Methods("POST")
	s.HandleFunc("/history/{id}", controllers.GetHistory).Methods("GET")
	s.HandleFunc("/history/user/{id}", controllers.GetAllHistory).Methods("GET")
	s.HandleFunc("/history/{id}", controllers.DeleteHistory).Methods("DELETE")

	s.HandleFunc("/popular/user/{id}", controllers.GetAllFavorite).Methods("GET")
	s.HandleFunc("/" +
		"ё/user/{id}", controllers.GetAllHistory).Methods("GET")

	return r
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
