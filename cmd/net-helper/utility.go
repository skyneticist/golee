package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

const colorRed = "\033[0;31m"
const colorGreen = "\033[0;32m"
const colorBlue = "\033[0;34m"
const colorNone = "\033[0m"

// RedString returns a red string.
func RedString(s string) string {
	return colorRed + s + colorNone
}

// GreenString returns a green string.
func GreenString(s string) string {
	return colorGreen + s + colorNone
}

// GreenString returns a green string.
func BlueString(s string) string {
	return colorBlue + s + colorNone
}

// // getGitBranch - Grabs current checkedout branch
// func getGitBranch() (string, error) {
// 	subCmd := []string{"rev-parse", "--abbrev-ref", "HEAD"}
// 	cmd := exec.Command("git", subCmd...)
// 	stdout, err := cmd.Output()
// 	if err != nil {
// 		return err.Error(), err
// 	}

// 	return string(stdout), nil
// }

func GetBranch() string {
	cmds := GitCmdList{
		GitCmd{
			cmd:  "branch",
			args: []string{"--show-current"},
		},
	}
	stdout, err := cmds.multipass()
	if err != nil {
		panic(err)
	}

	return string(stdout[0])
}

func CheckIfRemoteExists() bool {
	br := GetBranch()

	gitbranch := exec.Command("git", "branch")
	findstr := exec.Command("findstr", br)

	pipe, _ := gitbranch.StdoutPipe()
	defer pipe.Close()
	findstr.Stdin = pipe

	gitbranch.Start()

	remote, _ := findstr.Output()
	fmt.Println(remote)

	var remoteExists bool
	if string(remote) == "[]" {
		remoteExists = false
	} else {
		remoteExists = true
	}

	return remoteExists
}

func OpenPrompt() bool {
	fmt.Print("This is a fresh remote. Would you like to set upstream tracking? (y/n)")
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
