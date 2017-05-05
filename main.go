package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

func main() {
	var lang, sort, direction string
	app := cli.NewApp()
	app.Version = "1.0"
	app.Author = "monmaru"
	app.UsageText = "Usage: $ ghstar <GitHub User Name>"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Usage:       "Filter the repository by the name of the programming language",
			Destination: &lang,
		},
		cli.StringFlag{
			Name:        "sort, s",
			Usage:       "Sort. You can specify either created, updated or pushed",
			Value:       "created",
			Destination: &sort,
		},
		cli.StringFlag{
			Name:        "direction, d",
			Usage:       "Sorting direction. You can specify either desc or asc.",
			Value:       "desc",
			Destination: &direction,
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			fmt.Print(c.App.UsageText)
			return nil
		}

		params := &params{
			lang:      lang,
			sort:      sort,
			direction: direction,
		}
		return listRepositories(c.Args().First(), params)
	}
	app.Run(os.Args)
}

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

	for {
		starredRepos, res, err := client.Activity.ListStarred(context.Background(), user, opts)
		if err != nil {
			return err
		}

		for _, r := range starredRepos {
			repo := *r.Repository
			if !isEmpty(params.lang) && ((repo.Language == nil) || (*repo.Language != params.lang)) {
				continue
			}

			show(repo)
		}

		if res.NextPage == 0 {
			break
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

func show(repo github.Repository) {
	fmt.Println("-----------------------------------------")
	fmt.Println(*repo.HTMLURL)
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
