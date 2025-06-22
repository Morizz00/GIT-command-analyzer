package gitutils

import (
	"os/exec"
	"strings"
)

type Commit struct {
	Author string
	Date   string
	Hash   string
}

func GetCommit(repoPath string) ([]Commit, error) {
	cmd := exec.Command("git", "-C", repoPath, "log", "--pretty=format:%H|%an|%ad", "--date=short")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(out), "\n")
	var commits []Commit

	for _, line := range lines {
		parts := strings.Split(line, "|")

		if len(parts) == 3 {
			commits = append(commits, Commit{Hash: parts[0], Author: parts[1], Date: parts[2]})
		}
	}
	return commits, nil
}
