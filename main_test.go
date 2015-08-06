package main

import "testing"

func TestParseArgs(t *testing.T) {
	assert := func(args []string, user, repo string) {
		u, r, err := ParseArgs(args)
		if err != nil {
			t.Error(err)
		}
		if u != user {
			t.Errorf("User Expected: %s, but got %s", user, u)
		}
		if r != repo {
			t.Errorf("Repo Expected: %s, but got %s", repo, r)
		}
	}

	shouldErr := func(args []string) {
		_, _, err := ParseArgs(args)
		if err == nil {
			t.Error("Expected error, but got nil")
		}
	}

	assert([]string{"pocke/gh-issues"}, "pocke", "gh-issues")
	assert([]string{"pocke", "gh-issues"}, "pocke", "gh-issues")

	shouldErr([]string{})
	shouldErr([]string{"a", "b", "c"})
}
