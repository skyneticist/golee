package main

// GitCmd - Struct defining command and arguments for a specified GitCmd
type GitCmd struct {
	cmd  string
	args []string
}

// GitCmdList - Slice structure that holds one or more GitCmd structs
type GitCmdList []GitCmd
