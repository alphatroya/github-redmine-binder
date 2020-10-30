package main

import (
	"fmt"
	"testing"

	gh "github.com/google/go-github/v32/github"
)

func TestHighlightIssues(t *testing.T) {
	host := "https://google.com"
	fullHost := host + "/issues/"
	cases := []struct {
		input  string
		output string
	}{
		{
			"Test \n- #54444\n- #55555",
			fmt.Sprintf("Test \n- [#54444](%s54444)\n- [#55555](%s55555)", fullHost, fullHost),
		},
		{
			"Test \n- #54444\n- 55555",
			fmt.Sprintf("Test \n- [#54444](%s54444)\n- [#55555](%s55555)", fullHost, fullHost),
		},
		{
			"Test \n- #54444",
			fmt.Sprintf("Test \n- [#54444](%s54444)", fullHost),
		},
		{
			fmt.Sprintf("Test \n- [#54444](%s54444)", fullHost),
			fmt.Sprintf("Test \n- [#54444](%s54444)", fullHost),
		},
		{
			"Test \n- #54\n- #55",
			"Test \n- #54\n- #55",
		},
	}

	for _, testCase := range cases {
		handler := Handler{}
		pr := gh.PullRequest{Body: &testCase.input}
		result := handler.highlightLinks(&pr, host)
		if *result.Body != testCase.output {
			t.Errorf("Wrong hightlight result\nexpected: %s\nreceived: %s", testCase.output, *result.Body)
		}
	}
}
