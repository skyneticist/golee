package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	cli "github.com/urfave/cli/v2"
)

// multipass - Receiver function that takes a GitCmdList ([]string) and iterates
// over each entry, passing the entry to runGitCmd
func (gitCmds GitCmdList) multipass() ([]string, error) {
	if len(gitCmds) < 1 {
		panic("no arguments found! must have GitCmd entries!")
	}

	var result []string

	for i, gitCmd := range gitCmds {
		info, err := runGitCmd(gitCmd)
		if err != nil {
			return []string{"Error occurred in multipass function - line 13 in gitfuncs.go"}, err
		}
		if i == 0 {
			info = "\n" + info
			result = append(result, info)
		} else {
			result = append(result, info)
		}
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
		return "error at runGitCmd()!", err
	}

	return string(stdout), nil
}

// getGitBranch - Grabs current checkedout branch
func getGitBranch() (string, error) {
	subCmd := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	cmd := exec.Command("git", subCmd...)
	stdout, err := cmd.Output()
	if err != nil {
		return err.Error(), err
	}

	return string(stdout), nil
}

func Help(c *cli.Context) error {
	cli.ShowAppHelp(c)
	cli.ShowCommandHelp(c, "also-nope")
	cli.ShowSubcommandHelp(c)
	cli.ShowCompletions(c)
	cli.ShowCommandCompletions(c, "nope")

	c.Command.FullName()
	c.Command.HasName("gol")
	c.Command.Names()
	c.Command.VisibleFlags()

	return nil
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
		return err
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
			args: []string{"-u", "origin", "HEAD"},
		},
	}

	info, err := cmds.multipass()
	if err != nil {
		return err
	}
	fmt.Println(info)

	return nil
}

// StashPullPop - Runs git stash, pull, pop
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
		return err
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
		// GitCmd{
		// 	cmd:  "checkout",
		// 	args: []string{"main"},
		// },
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

// RenameBranch - Renames current branch to passed string
func RenameBranch(c *cli.Context) error {
	oldName, err := getGitBranch()
	oldName = ":" + oldName
	if err != nil {
		return err
	}

	newName := os.Args[2]
	cmds := GitCmdList{
		GitCmd{
			cmd:  "branch",
			args: []string{"-m", newName},
		},
		GitCmd{
			cmd:  "push",
			args: []string{"origin", oldName, newName},
		},
		GitCmd{
			cmd:  "push",
			args: []string{"origin", "-u", newName},
		},
	}

	info, err := cmds.multipass()
	if err != nil {
		return err
	}
	fmt.Println(info)

	return nil
}

func SetGitAuthors(c *cli.Context) error {
	argslength := len(os.Args[1:])
	if argslength < 2 {
		fmt.Println("need to pass at least one author.")
		return nil
	}

	var fmtAuthors string
	if argslength > 2 {
		fmtAuthors = os.Args[2] + ", " + os.Args[3]
	} else {
		fmtAuthors = os.Args[2]
	}

	cmds := GitCmdList{
		GitCmd{
			cmd:  "config",
			args: []string{"--global", "user.name", fmtAuthors},
		},
	}

	info, err := cmds.multipass()
	if err != nil {
		return err
	}

	fmt.Println(info)
	return nil
}

// CreateLogFile - Dumps git log into file
func CreateLogFile(c *cli.Context) error {
	cmds := GitCmdList{
		GitCmd{
			cmd:  "log",
			args: nil,
		},
	}

	info, err := cmds.multipass()
	if err != nil {
		return err
	}
	writeOutputToFile(info)

	return nil
}

func writeOutputToFile(data []string) {
	var bytes []byte
	for _, d := range data {
		bt := []byte(d)
		bytes = append(bytes, bt...)
	}
	home, _ := os.UserHomeDir()
	logLocation := filepath.Join(home, "git_log.txt")
	os.WriteFile(logLocation, bytes, 0644)
}
