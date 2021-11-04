package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	cli "github.com/urfave/cli/v2"
)

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
				Action:                 fullpull,
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
				Action:                 addCommitPush,
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
				Action:                 stashPullPop,
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
				Action:                 softReset,
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
				Action:                 hardReset,
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

func (gitCmds GitCmdList) multipass() error {
	if len(gitCmds) < 1 {
		panic("no arguments found! must have GitCmd entries!")
	}

	for _, pass := range gitCmds {
		result, err := runGitCmd(pass)
		if err != nil {
			return err
		}
		fmt.Println(result)
	}
	return nil
}

func runGitCmd(subCmd GitCmd) (string, error) {
	git := "git"
	args := []string{}

	args = append(args, subCmd.cmd)
	args = append(args, subCmd.args...)

	cmd := exec.Command(git, args...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return err.Error(), err
	}

	return string(stdout), nil
}

func fullpull(c *cli.Context) error {
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
	err := cmds.multipass()
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func addCommitPush(c *cli.Context) error {
	commitMsg := "msg goes here, dynamically"
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
			args: nil,
		},
	}
	err := cmds.multipass()
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func stashPullPop(c *cli.Context) error {
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
	err := cmds.multipass()
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func softReset(c *cli.Context) error {
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

	err := cmds.multipass()
	if err != nil {
		return err
	}
	return nil
}

func hardReset(c *cli.Context) error {
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

	err := cmds.multipass()
	if err != nil {
		return err
	}
	return nil
}
