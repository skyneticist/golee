package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

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

func checkIfRemoteExists() bool {
	ac := GitCmdList{
		GitCmd{
			cmd:  "branch",
			args: []string{"--show-current"},
		},
	}
	br, err := ac.multipass()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	cmds := GitCmdList{
		GitCmd{
			cmd:  "branch",
			args: []string{"--contains", string(br[0]), "|", "grep", "-w", br[0]},
		},
	}
	cmds.multipass()

	return false
}

func openprompt() bool {
	fmt.Println("This is a fresh remote. Would you like to set upstream tracking?")
	reader := bufio.NewReader(os.Stdin)
	rune, _, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}
	if rune == 'y' || rune == 'Y' {
		return true
	}
	if rune == 'n' || rune == 'Y' {
		return false
	}
	return false
}

// AddCommitPush - Add, Commit, Push local changes to current branch
func AddCommitPush(c *cli.Context) error {
	pushArgs := []string{}
	if !checkIfRemoteExists() {
		if openprompt() {
			pushArgs = append(pushArgs, []string{"-u", "origin", "HEAD"}...)
		}
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
			args: pushArgs,
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

// SoftReset - Undo staged changes while saving changes
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

// HardReset - Undo staged changes while destroying changes
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

// UndoMerge - Undo most recent merge or merge with give commit-id
func UndoMerge(c *cli.Context) error {
	var mergeId string
	if len(os.Args[1:]) > 0 {
		mergeId = os.Args[2]
	} else {
		mergeId = "HEAD"
	}
	cmds := GitCmdList{
		GitCmd{
			cmd:  "revert",
			args: []string{mergeId},
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

// SetGitAuthors - Sets authors for git commits
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

// CreateCheckoutBranch - Create then checkout branch with given name
func CreateCheckoutBranch(c *cli.Context) error {
	branchName := os.Args[2]
	cmds := GitCmdList{
		GitCmd{
			cmd:  "checkout",
			args: []string{"-b", branchName},
		},
	}
	info, err := cmds.multipass()
	if err != nil {
		return err
	}
	fmt.Println(info)

	return nil
}

// SearchLogPerMessage - Search the git log for a specific commit
// with the given commit message given
func SearchLogPerMessage(c *cli.Context) error {
	if len(os.Args) < 2 {
		fmt.Println("oof - you need to pass something to search on")
		return fmt.Errorf("error")
	}
	commitMsg := os.Args[2]
	cmds := GitCmdList{
		GitCmd{
			cmd:  "log",
			args: []string{"--pretty=oneline", "|", "grep", commitMsg},
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
	os.WriteFile("log_file.txt", bytes, 0644)
}
