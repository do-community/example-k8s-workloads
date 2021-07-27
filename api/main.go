package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "4000"
)

type User struct {
	Id                int    `json:"id"`
	UserName          string `json:"user_name"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	AvatarURL         string `json:"avatar_url"`
	Company           string `json:"company"`
	MostRecentComment string `json:"most_recent_comment"`
}

var users = []User{
	{Id: 654651651,
		UserName:          "kimschles",
		FirstName:         "Kim",
		LastName:          "Schlesinger",
		AvatarURL:         "https://community-cdn-digitalocean-com.global.ssl.fastly.net/variants/a79PCNZWVqKWL4pZZytLEXmW/1b33f0ae5d4693bf57c52014e04c03ab70f276df2ccd0b8ddde11732686ee1a9",
		Company:           "Digital Ocean",
		MostRecentComment: "Hey, I'm giving a tech talk on July 28, 2021. Join me!"},
	{Id: 654651652,
		UserName:          "masonegger",
		FirstName:         "Mason",
		LastName:          "Egger",
		AvatarURL:         "https://community-cdn-digitalocean-com.global.ssl.fastly.net/variants/EXJi5mGhdvFYbTxQ8Sfkdwfd/1b33f0ae5d4693bf57c52014e04c03ab70f276df2ccd0b8ddde11732686ee1a9",
		Company:           "Digital Ocean",
		MostRecentComment: "Python!",
	},
	{Id: 654651653,
		UserName:          "chrisoncode",
		FirstName:         "Chris",
		LastName:          "Sev",
		AvatarURL:         "https://community-cdn-digitalocean-com.global.ssl.fastly.net/variants/FvZ5kCncEfUQXbUkiSrNyJrW/1b33f0ae5d4693bf57c52014e04c03ab70f276df2ccd0b8ddde11732686ee1a9",
		Company:           "Digital Ocean",
		MostRecentComment: "JavaScript!",
	},
	{Id: 654651654,
		UserName:          "mattipv4",
		FirstName:         "Matt",
		LastName:          "Cowley",
		AvatarURL:         "https://community-cdn-digitalocean-com.global.ssl.fastly.net/variants/KyhncqSAFFeF3ULjGDX3MgC7/1b33f0ae5d4693bf57c52014e04c03ab70f276df2ccd0b8ddde11732686ee1a9",
		Company:           "Digital Ocean & CloudFlare",
		MostRecentComment: "Rails, CDNs, Community!",
	},
}

func writeJSONResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "s-maxage=3600, maxage=0")
	json.NewEncoder(w).Encode(v)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()

	showAllUserDataHandler := http.HandlerFunc(showAllUserData)
	mux.Handle("/", showAllUserDataHandler)

	showLatestCommentHandler := http.HandlerFunc(showLatestComment)
	mux.Handle("/comment", showLatestCommentHandler)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func showAllUserData(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	writeJSONResponse(w, users)

}

func showLatestComment(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	for _, user := range users {
		w.Write([]byte(fmt.Sprintf(user.FirstName + "'s Most recent comment: " + user.MostRecentComment + "\n")))
	}
}
