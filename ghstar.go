package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type params struct {
	lang, sort, direction string
}

func listRepositories(user string, params *params) error {
	client := newGitHubClient()
	opts := &github.ActivityListStarredOptions{
		Sort:      params.sort,
		Direction: params.direction,
	}
	opts.Page = 1
	opts.PerPage = 50

	for opts.Page != 0 {
		starredRepos, res, err := client.Activity.ListStarred(context.Background(), user, opts)
		if err != nil {
			return err
		}

		for _, r := range starredRepos {
			repo := r.Repository
			if isTargetLang(repo, params.lang) {
				show(repo)
			}
		}

		opts.Page = res.NextPage
	}
	return nil
}

func newGitHubClient() *github.Client {
	token := os.Getenv("GITHUB_API_TOKEN")
	if isEmpty(token) {
		return github.NewClient(nil)
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	return github.NewClient(oauth2.NewClient(context.Background(), ts))
}

func isTargetLang(r *github.Repository, lang string) bool {
	if isEmpty(lang) {
		return true
	}
	return (r.Language != nil) && (strings.ToLower(*r.Language) == strings.ToLower(lang))
}

func show(repo *github.Repository) {
	fmt.Println("-----------------------------------------")
	color.Green(*repo.HTMLURL)
	if repo.Description != nil {
		fmt.Println(*repo.Description)
	}
	if repo.Language != nil {
		fmt.Printf("[Lang:%v] ", *repo.Language)
	}
	if repo.StargazersCount != nil {
		fmt.Printf("[Star:%v] ", *repo.StargazersCount)
	}
	if repo.ForksCount != nil {
		fmt.Printf("[Fork:%v] ", *repo.ForksCount)
	}
	if repo.UpdatedAt != nil {
		fmt.Printf("[Updated at %v]", *repo.UpdatedAt)
	}
	fmt.Println("")
}

func isEmpty(s string) bool { return len(s) == 0 }
