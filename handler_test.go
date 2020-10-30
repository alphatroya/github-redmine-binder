package main

import (
	"fmt"
	gh "github.com/google/go-github/v32/github"
	"testing"
)

func TestHighlightIssues(t *testing.T) {
	host := "https://pm.handh.ru/issues/"
	cases := []struct{
		input string
		output string
	} {
		{
			"Test \n- #54444\n- #55555",
			fmt.Sprintf("Test \n- [#54444](%s54444)\n- [#55555](%s55555)", host, host),
		},
		{
			"Test \n- #54444\n- 55555",
			fmt.Sprintf("Test \n- [#54444](%s54444)\n- [#55555](%s55555)", host, host),
		},
		{
			"Test \n- #54444",
			fmt.Sprintf("Test \n- [#54444](%s54444)", host),
		},
		{
			fmt.Sprintf("Test \n- [#54444](%s54444)", host),
			fmt.Sprintf("Test \n- [#54444](%s54444)", host),
		},
		{
			"Test \n- #54\n- #55",
			"Test \n- #54\n- #55",
		},
	}

	for _, testCase := range cases {
		handler := Handler{}
		pr := gh.PullRequest{Body: &testCase.input}
		result := handler.highlightLinks(&pr)
		if *result.Body != testCase.output {
			t.Errorf("Wrong hightlight result\nexpected: %s\nreceived: %s", testCase.output, *result.Body)
		}
	}
}
