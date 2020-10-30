package main

import (
	"log"
	"net/http"
	"os"
)

const (
	accessTokenENV = "GITHUB_ACCESS_TOKEN"
)

func main() {
	accessToken := os.Getenv(accessTokenENV)
	if len(accessToken) == 0 {
		log.Fatalf("ENV %s is not set", accessTokenENV)
		return
	}
	http.Handle("/github", &Handler{githubAccessToken: accessToken})
	port := "8933"
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
