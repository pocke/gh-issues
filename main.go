package main

import (
	"fmt"

	"github.com/google/go-github/github"
)

func main() {
	c := github.NewClient(nil)
	opt := &github.IssueListByRepoOptions{}
	i, _, err := c.Issues.ListByRepo("pocke", "dotfiles", opt)
	if err != nil {
		panic(err)
	}

	for _, v := range i {
		fmt.Println(*v.Title)
	}
}
