package main

import (
	"fmt"
	"log"
	"os"
	"time"

	cli "github.com/urfave/cli/v2"
)

func main() {
	// initialize config for cli app
	InitApp()

	app := &cli.App{
		Name:        "git golee",
		HelpName:    "",
		Usage:       "git cli helper for productivity",
		UsageText:   "",
		ArgsUsage:   "",
		Version:     "0.1",
		Description: "",
		Commands: []*cli.Command{
			&FullPullCmd,
			&AddCpCmd,
			&AddCpRCmd,
			&StashPpCmd,
			&SoftResetCmd,
			&HardResetCmd,
			&UndoMergeCmd,
			&RenameBranchCmd,
			&CreateCheckoutBranchCmd,
			&SearchLogPerMessageCmd,
			&SetGitAuthorsCmd,
			&CreateLogFileCmd,
		},
		Flags:                []cli.Flag{},
		EnableBashCompletion: true,
		HideHelp:             false,
		HideHelpCommand:      false,
		HideVersion:          false,
		BashComplete: func(c *cli.Context) {
			fmt.Fprintf(c.App.Writer, "--better\n")
		},
		Action: func(c *cli.Context) error {
			fmt.Println()
			fmt.Println("Welcome to git-golee!")
			cli.ShowVersion(c)
			fmt.Println()

			// cli.DefaultAppComplete(c)
			fmt.Println("For usage info, run 'golee help'")
			return nil
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			fmt.Fprintf(c.App.Writer, "for shame\n")
			return err
		},
		Compiled: time.Time{},
		Authors: []*cli.Author{
			{Name: "skyneticist", Email: "hunterhartline@gmail.com"},
		},
		Copyright: "2021",
		Reader:    nil,
		Writer:    nil,
		ErrWriter: nil,
		ExitErrHandler: func(context *cli.Context, err error) {
		},
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
