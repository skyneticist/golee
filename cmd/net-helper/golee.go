package main

import (
	"fmt"
	"log"
	"os"
	"time"

	cli "github.com/urfave/cli/v2"
)

// TODO: add comments where it makes sense
// TODO: add --set-upstream workflow in ACP Action
// TODO: make output colorful like native git cmds
// TODO: add fx for appending the story number from branch name to beginning of commit

func main() {
	// initialize configuration
	// for cli app
	InitApp()

	app := &cli.App{
		Name:        "git golee",
		HelpName:    "",
		Usage:       "git cli helper for productivity",
		UsageText:   "",
		ArgsUsage:   "",
		Version:     "",
		Description: "",
		Commands: []*cli.Command{
			&FullPullCmd,
			&AddCpCmd,
			&StashPpCmd,
			&SoftResetCmd,
			&HardResetCmd,
			&UndoMergeCmd,
		},
		Flags:                []cli.Flag{},
		EnableBashCompletion: false,
		HideHelp:             false,
		HideHelpCommand:      false,
		HideVersion:          false,
		BashComplete: func(c *cli.Context) {
			fmt.Fprintf(c.App.Writer, "--better\n")
		},
		Action: func(c *cli.Context) error {
			c.Command.FullName()
			c.Command.HasName("gol")
			c.Command.Names()
			c.Command.VisibleFlags()

			return nil
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			fmt.Fprintf(c.App.Writer, "for shame\n")
			return err
		},
		Compiled: time.Time{},
		Authors: []*cli.Author{
			{Name: "skyneticist", Email: "hunterhartline@mail"},
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
