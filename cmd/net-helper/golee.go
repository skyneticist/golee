package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	cli "github.com/urfave/cli/v2"
)

// TODO: add comments where it makes sense
// TODO: make output colorful like native git cmds
// TODO: add fx for appending the story number from branch name to beginning of commit
// TODO: break out functions and cli.Commands to their own individual files (gitfuncs.go, commands.go)
// TODO: much more to come

type GitCmdList []GitCmd

type GitCmd struct {
	cmd  string
	args []string
}

func init() {
	cli.AppHelpTemplate += "\nThis might be helpful: \n"
	cli.CommandHelpTemplate += "\nThis is command help: \n"
	cli.SubcommandHelpTemplate += "\nThis is subcommand help"

	cli.HelpFlag = &cli.BoolFlag{Name: "help", Aliases: []string{"h", "H", "halp", "wtf"}}
	cli.VersionFlag = &cli.BoolFlag{Name: "print-version", Aliases: []string{"V", "v", "ver"}}

	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		fmt.Fprintf(w, "Never fear, HELP is here: ")
	}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "version=%s\n", c.App.Version)
	}
	cli.OsExiter = func(c int) {
		fmt.Fprintf(cli.ErrWriter, "refusing to exit %d\n", c)
	}

	cli.ErrWriter = ioutil.Discard

	cli.FlagStringer = func(fl cli.Flag) string {
		return fmt.Sprintf("\t\t%s", fl.Names()[0])
	}
}

func main() {
	app := &cli.App{
		Name:        "git golee",
		HelpName:    "",
		Usage:       "git cli helper for productivity",
		UsageText:   "",
		ArgsUsage:   "",
		Version:     "",
		Description: "",
		Commands: []*cli.Command{
			{
				Name:        "fullpull",
				Aliases:     []string{"fp"},
				Usage:       "safely pull down all changes for all projects in a given folder",
				UsageText:   "gg fp",
				Description: "automates pull on all repos in directory",
				ArgsUsage:   "",
				Category:    "",
				BashComplete: func(c *cli.Context) {
					fmt.Fprintf(c.App.Writer, "--better\n")
				},
				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					fmt.Fprintf(c.App.Writer, "for shame\n")
					return err
				},
				Action:                 Fullpull,
				Subcommands:            []*cli.Command{},
				Flags:                  []cli.Flag{},
				SkipFlagParsing:        false,
				HideHelp:               false,
				HideHelpCommand:        false,
				Hidden:                 false,
				UseShortOptionHandling: false,
				HelpName:               "",
				CustomHelpTemplate:     "",
			},
			{
				Name:        "addCommitPush",
				Aliases:     []string{"acp"},
				Usage:       "add -> commit -> push in one command",
				UsageText:   "gg acp",
				Description: "automates remote push of changes in current branch",
				ArgsUsage:   "",
				Category:    "",
				BashComplete: func(c *cli.Context) {
					fmt.Fprintf(c.App.Writer, "--better\n")
				},
				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					fmt.Fprintf(c.App.Writer, "for shame\n")
					return err
				},
				Action:                 AddCommitPush,
				Subcommands:            []*cli.Command{},
				Flags:                  []cli.Flag{},
				SkipFlagParsing:        false,
				HideHelp:               false,
				HideHelpCommand:        false,
				Hidden:                 false,
				UseShortOptionHandling: false,
				HelpName:               "",
				CustomHelpTemplate:     "",
			},
			{
				Name:        "stashPullPop",
				Aliases:     []string{"spp"},
				Usage:       "add -> commit -> push in one command",
				UsageText:   "gg spp",
				Description: "automates remote push of changes in current branch",
				ArgsUsage:   "",
				Category:    "",
				BashComplete: func(c *cli.Context) {
					fmt.Fprintf(c.App.Writer, "--better\n")
				},
				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					fmt.Fprintf(c.App.Writer, "for shame\n")
					return err
				},
				Action:                 StashPullPop,
				Subcommands:            []*cli.Command{},
				Flags:                  []cli.Flag{},
				SkipFlagParsing:        false,
				HideHelp:               false,
				HideHelpCommand:        false,
				Hidden:                 false,
				UseShortOptionHandling: false,
				HelpName:               "",
				CustomHelpTemplate:     "",
			},
			{
				Name:        "softReset",
				Aliases:     []string{"sr", "soft"},
				Usage:       "git soft reset commit in one command",
				UsageText:   "gg soft",
				Description: "uncommits most recent commit, keeps changes intact",
				ArgsUsage:   "",
				Category:    "",
				BashComplete: func(c *cli.Context) {
					fmt.Fprintf(c.App.Writer, "--better\n")
				},
				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					fmt.Fprintf(c.App.Writer, "for shame\n")
					return err
				},
				Action:                 SoftReset,
				Subcommands:            []*cli.Command{},
				Flags:                  []cli.Flag{},
				SkipFlagParsing:        false,
				HideHelp:               false,
				HideHelpCommand:        false,
				Hidden:                 false,
				UseShortOptionHandling: false,
				HelpName:               "",
				CustomHelpTemplate:     "",
			},
			{
				Name:        "hardReset",
				Aliases:     []string{"hr", "hard"},
				Usage:       "git hard reset commit in one command",
				UsageText:   "gg hard",
				Description: "uncommits most recent commit, destroys changes from commit as well!",
				ArgsUsage:   "",
				Category:    "",
				BashComplete: func(c *cli.Context) {
					fmt.Fprintf(c.App.Writer, "--better\n")
				},
				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					fmt.Fprintf(c.App.Writer, "for shame\n")
					return err
				},
				Action:                 HardReset,
				Subcommands:            []*cli.Command{},
				Flags:                  []cli.Flag{},
				SkipFlagParsing:        false,
				HideHelp:               false,
				HideHelpCommand:        false,
				Hidden:                 false,
				UseShortOptionHandling: false,
				HelpName:               "",
				CustomHelpTemplate:     "",
			},
			{
				Name:        "undoMerge",
				Aliases:     []string{"um", "undomerge"},
				Usage:       "undo last merge in one command",
				UsageText:   "gg um",
				Description: "reverts master branch to state before selected/last merge",
				ArgsUsage:   "",
				Category:    "",
				BashComplete: func(c *cli.Context) {
					fmt.Fprintf(c.App.Writer, "--better\n")
				},
				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					fmt.Fprintf(c.App.Writer, "for shame\n")
					return err
				},
				Action:                 UndoMerge,
				Subcommands:            []*cli.Command{},
				Flags:                  []cli.Flag{},
				SkipFlagParsing:        false,
				HideHelp:               false,
				HideHelpCommand:        false,
				Hidden:                 false,
				UseShortOptionHandling: false,
				HelpName:               "",
				CustomHelpTemplate:     "",
			},
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
