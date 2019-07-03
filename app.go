package main

import (
	"fmt"
	"regexp"
	"strings"

	"net/http"

	"github.com/andygrunwald/go-jira"
	"gopkg.in/go-playground/webhooks.v5/github"
)

var regexProjectKey = "\\[[A-Z]*\\-[0-9]+\\]"
var port = ":8080"

func main() {
	tp := jira.BasicAuthTransport{
		Username: "ramadhanm1998@gmail.com",
		Password: "icB26nXqVx90BRVTrxKKB68F",
	}

	client, _ := jira.NewClient(tp.Client(), "https://m-f-hafizh.atlassian.net/")

	hook, _ := github.New(github.Options.Secret("1234567890"))
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				fmt.Println("No event found")
				w.WriteHeader(http.StatusInternalServerError)
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
			reg, _ := regexp.Compile(regexProjectKey)
			title := pullRequest.PullRequest.Title
			issueKey := strings.Replace(strings.Replace(reg.FindString(title), "[", "", -1), "]", "", -1)
			issue, _, err := client.Issue.Get(issueKey, nil)
			transitions, _, err := client.Issue.GetTransitions(issueKey)
			if err != nil {
				fmt.Println("Error : ", err)
			}
			fmt.Println(issue.Fields.Summary)
			fmt.Println(transitions[1].ID)
			// Do whatever you want from here...
			if pullRequest.Action == "edited" {
				// res, err := client.Issue.DoTransition(issue.ID, transitions[1].ID)
				i := &jira.Issue{
					Key: issueKey,
					Fields: &jira.IssueFields{
						Description: "edit description 1",
					},
				}
				res, _, err := client.Issue.Update(i)
				if err != nil {
					fmt.Println("Error : ", err)
				}
				fmt.Println("response : ", res)
			}
			// fmt.Printf("%+v", pullRequest)
		}
	})
	fmt.Printf("Running on port %s\n", port)
	http.ListenAndServe(port, nil)
}
