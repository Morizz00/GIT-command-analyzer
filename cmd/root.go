package cmd

import (
	"log"
	"os"

	"git-commit-analyzer/analyzer"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "gca",
	Short: "Git Commit Analyzer-View stats on your git repo",
	Run: func(cmd *cobra.Command, args []string) {
		status, err := analyzer.Analyze(".")
		if err != nil {
			log.Fatal("Error:", err)
			os.Exit(1)
		}
		status.Print()
	},
}

func Execute() {
	if err := root.Execute(); err != nil {
		log.Fatalf("Command failed:%v", err)
	}
}
