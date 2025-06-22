package analyzer

import (
	"git-commit-analyzer/gitutils"
)

func Analyze(pathRepo string) (*Status, error) {
	commit, err := gitutils.GetCommit(pathRepo)
	if err != nil {
		return nil, err
	}
	stats := ComputeStats(commit)
	return stats, nil
}
