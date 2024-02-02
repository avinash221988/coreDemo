package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)))

	issue := getIssueFromPayload()
	labelsToAdd := []string{"bug", "enhancement"} // Add your desired labels here

	_, _, err := client.Issues.AddLabelsToIssue(ctx, "OWNER", "REPO", issue.GetNumber(), labelsToAdd)
	if err != nil {
		log.Fatalf("Error adding labels: %v", err)
	}
}

func getIssueFromPayload() *github.Issue {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	payload, err := os.ReadFile(eventPath)
	if err != nil {
		log.Fatalf("Error reading event payload: %v", err)
	}

	var issue github.Issue
	if err := github.UnmarshalJSON(payload, &issue); err != nil {
		log.Fatalf("Error unmarshalling issue payload: %v", err)
	}

	return &issue
}
