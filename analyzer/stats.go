package analyzer

import (
	"fmt"
	"git-commit-analyzer/formatter"
	"git-commit-analyzer/gitutils"
	"sort"
)

type Status struct {
	TotalCommits int
	ByAuthor     map[string]int
	ByDate       map[string]int
}

func ComputeStats(commits []gitutils.Commit) *Status {
	s := &Status{
		TotalCommits: len(commits),
		ByAuthor:     map[string]int{},
		ByDate:       map[string]int{},
	}
	for _, c := range commits {
		s.ByAuthor[c.Author]++
		s.ByDate[c.Date]++
	}
	return s
}
func (s *Status) Print() {
	fmt.Printf("total commits: %d\n\n", s.TotalCommits)
	formatter.PrintHeader("Commits by Author")
	type authorCnt struct {
		Name string
		Cnt  int
	}
	var authors []authorCnt
	for name, cnt := range s.ByAuthor {
		authors = append(authors, authorCnt{name, cnt})
	}
	sort.Slice(authors, func(i, j int) bool {
		return authors[i].Cnt > authors[j].Cnt
	})
	for _, auth := range authors {
		formatter.PrintValue("👤 "+auth.Name, auth.Cnt)
	}
	formatter.PrintHeader("Commits by Date")
	dates := make([]string, 0, len(s.ByDate))
	for date := range s.ByDate {
		dates = append(dates, date)
	}
	sort.Strings(dates)
	for _, dih := range dates {
		formatter.PrintValue("📅 "+dih, s.ByDate[dih])
	}
}
