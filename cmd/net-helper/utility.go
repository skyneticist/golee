package main

import (
	"bufio"
	"bytes"
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
	fmt.Println(br)

	branch := exec.Command("git", "branch")
	findBranch := exec.Command("findstr", br)

	findBranch.Stdin, _ = branch.StdoutPipe()
	findBranch.Stdout = os.Stdout

	var buff bytes.Buffer
	_ = findBranch.Start()
	_ = branch.Run()
	_ = findBranch.Wait()
	findBranch.Stdout = &buff
	fmt.Println(buff.String())
	// gitbranch := exec.Command("git", "branch")
	// findstr := exec.Command("findstr", br)

	// reader, writer := io.Pipe()

	// gitbranch.Stdout = writer
	// findstr.Stdin = reader

	// var buff bytes.Buffer
	// findstr.Stdout = &buff

	// gitbranch.Start()
	// gitbranch.Wait()
	// findstr.Start()
	// findstr.Wait()
	// writer.Close()

	// total := buff.String()

	// fmt.Printf("Total processes running : %s", total)

	var remoteExists bool
	if len(buff.String()) == 0 {
		remoteExists = false
	} else {
		remoteExists = true
	}

	return remoteExists
}

// OpenPrompt opens cli prompt asking if the new branch should be pushed with upstream tracking set
func OpenPrompt() bool {
	fmt.Print(GreenString("This is a fresh remote. Would you like to set upstream tracking? (y/n)"))
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
