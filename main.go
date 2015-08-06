package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
)

func ParseArgs(args []string) (user, repo string, err error) {
	e := fmt.Errorf("Expected `user repo` or `user/repo`")
	switch len(args) {
	case 2:
		return args[0], args[1], nil
	case 1:
		a := strings.Split(args[0], "/")
		if len(a) != 2 {
			return "", "", e
		}
		return a[0], a[1], nil
	default:
		return "", "", e
	}
}

// mockable
var execGit = func() ([]byte, error) {
	return exec.Command("git", "remote", "-v").Output()
}

func DetectUserRepoFromGit() (user, repo string, err error) {
	o, err := execGit()
	if err != nil {
		return "", "", err
	}
	s := strings.Split(string(o), "\n")[0]
	re := regexp.MustCompile(`origin\s+.+github\.com.+?([[:alnum:]][[:alnum:]-]*)/([[:alnum:]._-]+?)(?:\.git)?\s+\(.+\)$`)
	ma := re.FindStringSubmatch(s)
	user, repo = ma[1], ma[2]
	if user == "" || repo == "" {
		return "", "", fmt.Errorf("Cannot detect user/repo from git")
	}
	return user, repo, nil
}

func E(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	user, repo, err := ParseArgs(os.Args[1:])
	if err != nil {
		user, repo, err = DetectUserRepoFromGit()
		if err != nil {
			E(err)
		}
	}

	c := github.NewClient(nil)
	opt := &github.IssueListByRepoOptions{}
	i, _, err := c.Issues.ListByRepo(user, repo, opt)
	if err != nil {
		E(err)
	}

	for _, v := range i {
		fmt.Printf("%d %s\n", *v.Number, *v.Title)
	}
}
