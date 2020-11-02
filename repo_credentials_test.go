package main

import "testing"

func TestExtractCredentialsCorrect(t *testing.T) {
	tcs := []struct{
		url string
		owner string
		repo string
	}{
		{
			url: "https://github.com/alphatroya/github-redmine-binder/pull/120",
			owner: "alphatroya",
			repo: "github-redmine-binder",
		},
		{
			url: "https://github.com/owef/repo/pull/120",
			owner: "owef",
			repo: "repo",
		},
	}
	for _, tc := range tcs {
		data, err := extractCredentials(tc.url)
		if err != nil {
			t.Fatalf("received non-ecpected error from url: %s\nerror: %s", tc.url, err)
		}
		if tc.owner != data.owner {
			t.Fatalf("wrong owner received from url: %s\nreceived: %s\nexpected: %s", tc.url, tc.owner, data.owner)
		}
		if tc.repo != data.repo {
			t.Fatalf("wrong repo received from url: %s\nreceived: %s\nexpected: %s", tc.url, tc.repo, data.repo)
		}
	}
}

func TestExtractCredentialsNonCorrect(t *testing.T) {
	tcs := []struct{
		url string
	}{
		{
			url: "https://github.com/owef/repo",
		},
		{
			url: "/github.com/owef/repo/pull/120",
		},
	}
	for _, tc := range tcs {
		_, err := extractCredentials(tc.url)
		if err == nil {
			t.Fatalf("received nil error from url: %s", tc.url)
		}
	}
}
