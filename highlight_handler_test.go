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
			fmt.Sprintf("Test [#54444](%s54444), [#54445](%s54445)", fullHost, fullHost),
			fmt.Sprintf("Test [#54444](%s54444), [#54445](%s54445)", fullHost, fullHost),
		},
		{
			fmt.Sprintf("Test \n- [#54444](%s54444)\n- [#54445](%s54445)", fullHost, fullHost),
			fmt.Sprintf("Test \n- [#54444](%s54444)\n- [#54445](%s54445)", fullHost, fullHost),
		},
		{
			fmt.Sprintf("Test \n- [#54444](%s54444)", fullHost),
			fmt.Sprintf("Test \n- [#54444](%s54444)", fullHost),
		},
		{
			"Test \n- #54\n- #55",
			"Test \n- #54\n- #55",
		},
		{
			"Test 54453",
			fmt.Sprintf("Test [#54453](%s54453)", fullHost),
		},
		{
			"#54453",
			fmt.Sprintf("[#54453](%s54453)", fullHost),
		},
		{
			"Test #54453",
			fmt.Sprintf("Test [#54453](%s54453)", fullHost),
		},
		{
			"Test #54453 #55743",
			fmt.Sprintf("Test [#54453](%s54453) [#55743](%s55743)", fullHost, fullHost),
		},
		{
			fmt.Sprintf("%s54453", fullHost),
			fmt.Sprintf("[#54453](%s54453)", fullHost),
		},
	}

	for _, testCase := range cases {
		handler := HighlightHandler{}
		pr := gh.PullRequest{Body: &testCase.input}
		result := handler.highlightLinks(&pr, host)
		if *result.Body != testCase.output {
			t.Errorf("Wrong hightlight result\nexpected: %s\nreceived: %s", testCase.output, *result.Body)
		}
	}
}
