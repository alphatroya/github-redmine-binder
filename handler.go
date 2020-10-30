package main

import (
	"context"
	"net/http"
	"regexp"

	gh "github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
	githook "gopkg.in/go-playground/webhooks.v5/github"
)

type Handler struct {
	githubAccessToken string
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	hook, _ := githook.New()
	payload, err := hook.Parse(request, githook.PullRequestEvent)
	if err != nil {
		return
	}

	switch payload.(type) {
	case githook.PullRequestPayload:
		pr := payload.(githook.PullRequestPayload)
		err = h.handlePR(pr)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(err.Error()))
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func (h Handler) handlePR(pr githook.PullRequestPayload) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: h.githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := gh.NewClient(tc)
	// TODO: parse input or move to ENV
	owner := "Heads-and-Hands"
	repo := "citilink-ios"
	number := int(pr.PullRequest.Number)
	sourcePR, _, err := client.PullRequests.Get(ctx, owner, repo, number)
	if err != nil {
		return err
	}
	editedPR := h.highlightLinks(sourcePR)
	_, _, err = client.PullRequests.Edit(ctx, owner, repo, number, editedPR)
	if err != nil {
		return err
	}
	return nil
}

func (h Handler) highlightLinks(pr *gh.PullRequest) *gh.PullRequest {
	body := pr.Body
	re := regexp.MustCompile(`- #?([0-9]{4,})`)
	replaced := re.ReplaceAllString(*body, "- [#$1](https://pm.handh.ru/issues/$1)")
	resultPR := gh.PullRequest{Body: &replaced}
	return &resultPR
}
