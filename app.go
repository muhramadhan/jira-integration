package main

import (
	"fmt"

	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

func main() {
	// tp := jira.BasicAuthTransport{
	// 	Username: "ramadhanm1998@gmail.com",
	// 	Password: "icB26nXqVx90BRVTrxKKB68F",
	// }

	// client, _ := jira.NewClient(tp.Client(), "https://m-f-hafizh.atlassian.net/")

	hook, _ := github.New(github.Options.Secret("1234567890"))
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				fmt.Println("No event found")
				return
			}
		}
		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})
}
