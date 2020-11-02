package main

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	gh "github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
	githook "gopkg.in/go-playground/webhooks.v5/github"
)

type HighlightHandler struct {
	githubAccessToken string
	redmineHost       string
}

func (h HighlightHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

func (h HighlightHandler) handlePR(pr githook.PullRequestPayload) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: h.githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := gh.NewClient(tc)
	prData, err := extractCredentials(pr.PullRequest.HTMLURL)
	if err != nil {
		return err
	}
	number := int(pr.PullRequest.Number)
	sourcePR, _, err := client.PullRequests.Get(ctx, prData.owner, prData.repo, number)
	if err != nil {
		return err
	}
	editedPR := h.highlightLinks(sourcePR, h.redmineHost)
	_, _, err = client.PullRequests.Edit(ctx, prData.owner, prData.repo, number, editedPR)
	if err != nil {
		return err
	}
	return nil
}

func (h HighlightHandler) highlightLinks(pr *gh.PullRequest, host string) *gh.PullRequest {
	body := pr.Body
	re := regexp.MustCompile(`- #?([0-9]{4,})`)
	replaced := re.ReplaceAllString(*body, fmt.Sprintf("- [#$1](%s/issues/$1)", host))
	resultPR := gh.PullRequest{Body: &replaced}
	return &resultPR
}
