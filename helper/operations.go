package helper

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/oliverhohn/many2mono/model"
)

func FetchRemote(r *model.Repo, dryRun bool) {
	cmd := exec.Command("git", "remote", "add", "-f", r.Name(), r.URL.String())

	err := runCommand(cmd, dryRun)
	if err != nil {
		log.Fatalf("unable to fetch remote %s, due to: %v", r.URL.String(), err)
	}
}

func RemoveRemote(r *model.Repo, dryRun bool) {
	cmd := exec.Command("git", "remote", "remove", r.Name())

	err := runCommand(cmd, dryRun)
	if err != nil {
		log.Fatalf("unable to remove remote %s, due to: %v", r.Name(), err)
	}
}

func MergeHistories(r *model.Repo, branch string, dryRun bool) {
	remote := fmt.Sprintf("%s/%s", r.Name(), branch)
	cmd := exec.Command("git", "merge", "--allow-unrelated-histories", "--strategy=ours", "--no-commit", remote)

	err := runCommand(cmd, dryRun)
	if err != nil {
		log.Fatalf("unable to merge histories from remote %s, due to: %v", remote, err)
	}
}

func PrefixFiles(r *model.Repo, branch string, dryRun bool) {
	remote := fmt.Sprintf("%s/%s", r.Name(), branch)
	cmd := exec.Command("git", "read-tree", fmt.Sprintf("--prefix=%s", r.NameWithoutOrg()), "-u", remote)

	err := runCommand(cmd, dryRun)
	if err != nil {
		log.Fatalf("unable to prefix files from remote %s, due to: %v", remote, err)
	}
}

func CommitChange(r *model.Repo, dryRun bool) {
	cmd := exec.Command("git", "commit", fmt.Sprintf("--message=\"Merged %s\"", r.Name()))

	err := runCommand(cmd, dryRun)
	if err != nil {
		log.Fatalf("unable to commit merge from remote %s, due to: %v", r.Name(), err)
	}
}

func runCommand(cmd *exec.Cmd, dryRun bool) error {
	fmt.Printf("Running %s\n", cmd.String())

	if !dryRun {
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("%v\ntrace:%s", err, out)
		}
	}

	return nil
}
