package main

import (
	"bufio"
	"fmt"
	"log"
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

	out, err := exec.Command("git", "branch", "-l", string(br)).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

	if string(out) != "" {
		return true
	} else {
		return false
	}
}

func CheckIfRemoteEx() bool {
	br := GetBranch()
	fmt.Println(br)

	out, err := exec.Command("git", "branch").Output()
	if err != nil {
		log.Fatal("git branch did not run successfully")
	}
	fmt.Printf("Output from branch: %s", out)

	nextOut := exec.Command("findstr", "main")
	if err != nil {
		log.Fatal("findstr did not run successfully")
	}

	pipe, err := nextOut.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	pipe.Write(out)

	nextOut.Start()
	finalerr := nextOut.Wait()
	if finalerr != nil {
		fmt.Println(err)
	}

	output, err := nextOut.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final output: %s", output)
	fmt.Printf("Output from findstr: %s", nextOut)

	// var remoteExists bool
	// if len(stdout) == 0 {
	// 	remoteExists = false
	// } else {
	// 	remoteExists = true
	// }

	return true
	// return remoteExists
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
