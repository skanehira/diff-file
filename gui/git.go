package gui

import (
	"os/exec"
	"strings"
)

type Commit struct {
	Hash    string
	Author  string
	Message string
}

func Commits(target string) ([]Commit, error) {
	args := []string{
		"log",
		`--pretty=format:%h %an %s`,
		target,
	}

	cmd := exec.Command("git", args...)
	res, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	output := strings.Split(string(res), "\n")
	return parse2Commit(output), nil
}

func parse2Commit(output []string) []Commit {
	var commits = make([]Commit, 0, len(output))
	for _, s := range output {
		l := strings.SplitN(s, " ", 3)
		if len(l) < 2 {
			// invalid format
			break
		}
		commits = append(commits, Commit{
			Hash:    l[0],
			Author:  l[1],
			Message: l[2],
		})
	}

	return commits
}
