package main

import (
	"fmt"
	"strings"
)

type repoCredentials struct {
	owner string
	repo  string
}

func extractCredentials(url string) (*repoCredentials, error) {
	trimmed := strings.TrimPrefix(url, "https://github.com/")
	parts := strings.Split(trimmed, "/")
	if len(parts) != 4 {
		return nil, fmt.Errorf("wrong PR link string %s", url)
	}
	return &repoCredentials{owner: parts[0], repo: parts[1]}, nil
}
