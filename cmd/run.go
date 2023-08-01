package cmd

import (
	"log"
	"strings"

	"github.com/oliverhohn/many2mono/helper"
	"github.com/oliverhohn/many2mono/model"
	"github.com/spf13/cobra"
)

// CLI flags
var defaultBranch string
var dryRun bool

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "TBD",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		repos := []*model.Repo{}
		for _, path := range args {
			r, err := model.NewRepo(path)
			if err != nil {
				log.Fatal(err)
			}

			repos = append(repos, r)
		}

		if duplicateRepos := findDuplicateRepoNames(repos); len(duplicateRepos) > 0 {
			log.Fatalf("Cannot merge as repos with duplicate names found: %s", strings.Join(duplicateRepos, ", "))
		}

		for _, r := range repos {
			helper.FetchRemote(r, dryRun)
			helper.MergeHistories(r, defaultBranch, dryRun)
			helper.PrefixFiles(r, defaultBranch, dryRun)
			helper.CommitChange(r, dryRun)
			helper.RemoveRemote(r, dryRun)
		}
	},
}

func findDuplicateRepoNames(repos []*model.Repo) []string {
	reposByName := map[string][]*model.Repo{}

	for _, r := range repos {
		if _, ok := reposByName[r.NameWithoutOrg()]; !ok {
			reposByName[r.NameWithoutOrg()] = []*model.Repo{}
		}

		reposByName[r.NameWithoutOrg()] = append(reposByName[r.NameWithoutOrg()], r)
	}

	ret := []string{}
	for name, repos := range reposByName {
		if len(repos) > 1 {
			ret = append(ret, name)
		}
	}

	return ret
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&defaultBranch, "branch", "b", "main", "TBD")

	runCmd.Flags().BoolVar(&dryRun, "dry-run", false, "TBD")
}
