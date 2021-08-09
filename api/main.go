package main

//to do
// break out Users to another file
// make mostrecentcomment return json
// user/{username} in another branch
// license
// readme

import (
	"encoding/json"
	"fmt"
	"github/example-k8s-workloads/api/data"
	"log"
	"net/http"
	"os"

	"github.com/do-community/example-k8s-workloads/api/data"
	"github.com/gorilla/mux"
)

const (
	defaultPort = "4000"
)

	users := data.User{
	{Id: 654651651,
		UserName:          "kimschles",
		FirstName:         "Kim",
		LastName:          "Schlesinger",
		AvatarURL:         "https://community-cdn-digitalocean-com.global.ssl.fastly.net/variants/a79PCNZWVqKWL4pZZytLEXmW/1b33f0ae5d4693bf57c52014e04c03ab70f276df2ccd0b8ddde11732686ee1a9",
		Company:           "Digital Ocean",
		MostRecentComment: "Hey, I'm giving a tech talk on July 28, 2021. Join me!",
		LastLogin:         "Tue, 27 Jul 2021 21:39:08 +0000",
	},
	{Id: 654651652,
		UserName:          "masonegger",
		FirstName:         "Mason",
		LastName:          "Egger",
		AvatarURL:         "https://community-cdn-digitalocean-com.global.ssl.fastly.net/variants/EXJi5mGhdvFYbTxQ8Sfkdwfd/1b33f0ae5d4693bf57c52014e04c03ab70f276df2ccd0b8ddde11732686ee1a9",
		Company:           "Digital Ocean",
		MostRecentComment: "Python!",
		LastLogin:         "Mon, 24 May 2021 21:39:08 +0000",
	},
	{Id: 654651653,
		UserName:          "chrisoncode",
		FirstName:         "Chris",
		LastName:          "Sev",
		AvatarURL:         "https://community-cdn-digitalocean-com.global.ssl.fastly.net/variants/FvZ5kCncEfUQXbUkiSrNyJrW/1b33f0ae5d4693bf57c52014e04c03ab70f276df2ccd0b8ddde11732686ee1a9",
		Company:           "Digital Ocean",
		MostRecentComment: "JavaScript!",
		LastLogin:         "Sun, 25 Jul 2021 16:33:08 +0000",
	},
	{Id: 654651654,
		UserName:          "mattipv4",
		FirstName:         "Matt",
		LastName:          "Cowley",
		AvatarURL:         "https://community-cdn-digitalocean-com.global.ssl.fastly.net/variants/KyhncqSAFFeF3ULjGDX3MgC7/1b33f0ae5d4693bf57c52014e04c03ab70f276df2ccd0b8ddde11732686ee1a9",
		Company:           "Digital Ocean & CloudFlare",
		MostRecentComment: "Rails, CDNs, Community!",
		LastLogin:         "Wed, 28 Jul 2021 10:05:29 +0000",
	},
}

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
	w.WriteHeader(200)
	writeJSONResponse(w, users)
}

// find single user data
// return single user data(?)
func returnSingleUserData(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	var returnValue = ""

	for _, user := range users {
		if username == user.UserName {
			returnValue = user.UserName
		}
	}
	if returnValue != "" {
		writeJSONResponse(w, returnValue)
	} else {
		writeJSONError(w, 404)
	}
}

func returnLatestComment(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	for _, user := range users {
		//return as json
		w.Write([]byte(fmt.Sprintf(user.FirstName + "'s Most recent comment: " + user.MostRecentComment + "\n")))
	}
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	writeJSONError(w, http.StatusNotFound)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := mux.NewRouter()

	returnAllUserDataHandler := http.HandlerFunc(returnAllUserData)
	mux.Handle("/users", returnAllUserDataHandler)

	returnSingleUserDataHandler := http.HandlerFunc(returnSingleUserData)
	mux.Handle("/user/{username}", returnSingleUserDataHandler)

	returnLatestCommentHandler := http.HandlerFunc(returnLatestComment)
	mux.Handle("/comments", returnLatestCommentHandler)

	healthCheckHandler := http.HandlerFunc(healthCheck)
	mux.Handle("/health", healthCheckHandler)

	notFoundHandler := http.HandlerFunc(notFound)
	mux.Handle("/", notFoundHandler)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
