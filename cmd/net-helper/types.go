package main

/*
	GitCmdList

	- slice containing GitCmd(s)
*/
type GitCmdList []GitCmd

/*
	GitCmd

	- struct defining cmd and args for a specified GitCmd
*/
type GitCmd struct {
	cmd  string
	args []string
}
