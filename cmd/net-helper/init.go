package main

import (
	"fmt"
	"io"
	"io/ioutil"

	cli "github.com/urfave/cli/v2"
)

// InitApp - Setup cli app configuration
func InitApp() {
	cli.AppHelpTemplate += "\nThis might be helpful: \n"
	cli.CommandHelpTemplate += "\nThis is command help: \n"
	cli.SubcommandHelpTemplate += "\nThis is subcommand help"

	cli.HelpFlag = &cli.BoolFlag{Name: "help", Aliases: []string{"h", "H", "wtf"}}
	cli.VersionFlag = &cli.BoolFlag{Name: "print-version", Aliases: []string{"V", "v", "ver"}}

	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		fmt.Fprintf(w, "\nNever fear, HELP is here 🚀\n \nAdd, Commit, Push in one command:           golee acp [commit message here]\n\nAdd, Commit, Push a new remote branch:      golee acpr [commit message is here]\n\nChange git authors for pairing:             golee auth [author1] [author2]\n\nUndo last commit, preserving changes:       golee soft\n\nUndo last commit, destroying changes:       golee hard\n\nStash, Pull, Pop:                           golee spp\n\nRename current branch:                      golee rn [new_branch_name]\n\n")
	}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "v%s\n", c.App.Version)
	}
	cli.OsExiter = func(c int) {
		fmt.Fprintf(cli.ErrWriter, "refusing to exit %d\n", c)
	}

	cli.ErrWriter = ioutil.Discard

	cli.FlagStringer = func(fl cli.Flag) string {
		return fmt.Sprintf("\t\t%s", fl.Names()[0])
	}
}
