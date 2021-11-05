package main

import (
	"fmt"
	"os"
	"os/exec"

	cli "github.com/urfave/cli/v2"
)

// multipass - Receiver function that takes a GitCmdList ([]string) and iterates
// over each entry, passing the entry to runGitCmd
func (gitCmds GitCmdList) multipass() (string, error) {
	if len(gitCmds) < 1 {
		panic("no arguments found! must have GitCmd entries!")
	}

	var result string

	for _, pass := range gitCmds {
		res, err := runGitCmd(pass)
		if err != nil {
			return "error occurred at multipass()!", err
		}
		result = res
	}
	return result, nil
}

// runGitCmd - Takes GitCmd parameter, executing the git command passed
// then returning command output to stdout
func runGitCmd(subCmd GitCmd) (string, error) {
	git := "git"
	args := []string{}

	args = append(args, subCmd.cmd)
	args = append(args, subCmd.args...)

	cmd := exec.Command(git, args...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return "error at runGitCmd() !!!", err
	}

	return string(stdout), nil
}

// getGitBranch - Grabs current checkedout branch
func getGitBranch() (string, error) {
	cmd := exec.Command("git", "branch")
	stdout, err := cmd.Output()
	if err != nil {
		return err.Error(), err
	}

	return string(stdout), nil
}

// Fullpull - Stash local changes then pull remote changes
func Fullpull(c *cli.Context) error {
	cmds := GitCmdList{
		GitCmd{
			cmd:  "stash",
			args: nil,
		},
		GitCmd{
			cmd:  "checkout",
			args: []string{"main"},
		},
		GitCmd{
			cmd:  "pull",
			args: nil,
		},
	}

	info, err := cmds.multipass()

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(info)

	return nil
}

// AddCommitPush - Add, Commit, Push local changes to current branch
func AddCommitPush(c *cli.Context) error {
	commitMsg := os.Args[2]
	cmds := GitCmdList{
		GitCmd{
			cmd:  "add",
			args: []string{"."},
		},
		GitCmd{
			cmd:  "commit",
			args: []string{"-m", commitMsg},
		},
		GitCmd{
			cmd:  "push",
			args: nil,
		},
	}

	info, err := cmds.multipass()

	if err != nil {
		return err
	}
	fmt.Println(info)

	return nil
}

// AddCommitPushRemote - Add, Commit, Push local changes
// on fresh branch (sets upstream)
func AddCommitPushRemote(c *cli.Context) error {
	gitBranch, err := getGitBranch()
	if err != nil {
		return err
	}
	commitMsg := os.Args[2]
	cmds := GitCmdList{
		GitCmd{
			cmd:  "add",
			args: []string{"."},
		},
		GitCmd{
			cmd:  "commit",
			args: []string{"-m", commitMsg},
		},
		GitCmd{
			cmd:  "push",
			args: []string{"--set-upstream", "origin", gitBranch},
		},
	}

	info, err := cmds.multipass()
	if err != nil {
		return err
	}
	fmt.Println(info)

	return nil
}

// StashPullPop - Runs git stash, pull, pop.
func StashPullPop(c *cli.Context) error {
	cmds := GitCmdList{
		GitCmd{
			cmd:  "stash",
			args: nil,
		},
		GitCmd{
			cmd:  "pull",
			args: nil,
		},
		GitCmd{
			cmd:  "stash",
			args: []string{"pop"},
		},
	}

	info, err := cmds.multipass()

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(info)

	return nil
}

// SoftReset - Undo last commit while saving changes
func SoftReset(c *cli.Context) error {
	cmds := GitCmdList{
		GitCmd{
			cmd:  "reset",
			args: []string{"--soft", "HEAD^"},
		},
		GitCmd{
			cmd:  "status",
			args: nil,
		},
	}

	info, err := cmds.multipass()

	if err != nil {
		return err
	}
	fmt.Println(info)

	return nil
}

// HardReset - Undo last commit while destroying changes
func HardReset(c *cli.Context) error {
	cmds := GitCmdList{
		GitCmd{
			cmd:  "reset",
			args: []string{"--hard", "HEAD^"},
		},
		GitCmd{
			cmd:  "status",
			args: nil,
		},
	}

	info, err := cmds.multipass()

	if err != nil {
		return err
	}
	fmt.Println(info)

	return nil
}

// UndoMerge - Undo most recent merge unless commit hash passed
func UndoMerge(c *cli.Context) error {
	cmds := GitCmdList{
		GitCmd{
			cmd:  "checkout",
			args: []string{"main"},
		},
		GitCmd{
			cmd:  "log",
			args: []string{"--oneline"},
		},
	}

	info, err := cmds.multipass()
	if err != nil {
		return err
	}
	fmt.Println(info)

	return nil
}
