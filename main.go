package main

import (
	"log"
	"net/http"
	"os"
)

const (
	accessTokenENV = "GITHUB_ACCESS_TOKEN"
	redmineHostENV = "REDMINE_HOST"
)

func main() {
	accessToken := os.Getenv(accessTokenENV)
	if len(accessToken) == 0 {
		log.Fatalf("ENV %s is not set", accessTokenENV)
		return
	}
	host := os.Getenv(redmineHostENV)
	if len(host) == 0 {
		log.Fatalf("ENV %s is not set", host)
		return
	}
	http.Handle("/github", &Handler{githubAccessToken: accessToken, redmineHost: host})
	port := "8933"
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
