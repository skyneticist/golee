package main

import (
	"fmt"
	"os"
	"os/exec"

	cli "github.com/urfave/cli/v2"
)

/*
	multipass()

	- helper function
	- takes a slice of GitCmd(s) as only parameter
	- iterates over slice, passing entries to runGitCmd()
*/
func (gitCmds GitCmdList) multipass() (string, error) {
	if len(gitCmds) < 1 {
		panic("no arguments found! must have GitCmd entries!")
	}

	var result string

	for _, pass := range gitCmds {
		res, err := runGitCmd(pass)
		if err != nil {
			return err.Error(), err
		}
		result = res
	}
	return result, nil
}

/*
	runGitCmd()

	- main cmd which executes git cmds
	- takes a single GitCmd struct as parameter
	- executes the GitCmd that is passed
	- handles parsing of all subCmds
*/
func runGitCmd(subCmd GitCmd) (string, error) {
	git := "git"
	args := []string{}

	args = append(args, subCmd.cmd)
	args = append(args, subCmd.args...)

	cmd := exec.Command(git, args...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return err.Error(), err
	}

	return string(stdout), nil
}

/*
	FullPull()

	- Git Function
	- stashes local changes before checking out and pulling down main
	TODO: need to be able to pass another arg/flag to specify which branch (default: main)
*/
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
		fmt.Println(err.Error())
	}
	fmt.Println(info)

	return nil
}

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
