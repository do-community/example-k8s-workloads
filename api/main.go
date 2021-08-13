package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	defaultPort = "4000"
)

var (
	//go:embed users/users.json
	data []byte
)

type User struct {
	Id                int    `json:"id"`
	UserName          string `json:"user_name"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	AvatarURL         string `json:"avatar_url"`
	Company           string `json:"company"`
	MostRecentComment string `json:"most_recent_comment"`
	LastLogin         string `json:"last_login"`
}

var users []User

func writeJSONResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "s-maxage=3600, maxage=0")
	json.NewEncoder(w).Encode(v)
}

func writeJSONError(w http.ResponseWriter, code int) {
	msg := map[string]string{
		"error": http.StatusText(code),
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func returnAllUserData(w http.ResponseWriter, req *http.Request) {
	writeJSONResponse(w, users)
}

func returnSingleUserData(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]

	for _, user := range users {
		if username == user.UserName {
			writeJSONResponse(w, &user)
			return
		}
	}
	writeJSONError(w, 404)
}

type Comments struct {
	UserName          string `json:"user_name"`
	MostRecentComment string `json: "most_recent_comment"`
}

func returnLatestComment(w http.ResponseWriter, req *http.Request) {
	var recentComments []Comments
	for _, user := range users {
		recentComments = append(recentComments, Comments{
			UserName:          user.UserName,
			MostRecentComment: user.MostRecentComment,
		})
	}
	writeJSONResponse(w, recentComments)
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	writeJSONError(w, http.StatusNotFound)
}

func main() {

	err := json.Unmarshal([]byte(data), &users)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := mux.NewRouter()
	subrouter := mux.PathPrefix("/api/v1").Subrouter()

	returnAllUserDataHandler := http.HandlerFunc(returnAllUserData)
	subrouter.Handle("/users", returnAllUserDataHandler)

	returnSingleUserDataHandler := http.HandlerFunc(returnSingleUserData)
	subrouter.Handle("/user/{username}", returnSingleUserDataHandler)

	returnLatestCommentHandler := http.HandlerFunc(returnLatestComment)
	subrouter.Handle("/comments", returnLatestCommentHandler)

	healthCheckHandler := http.HandlerFunc(healthCheck)
	subrouter.Handle("/health", healthCheckHandler)

	notFoundHandler := http.HandlerFunc(notFound)
	subrouter.Handle("/", notFoundHandler)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}
