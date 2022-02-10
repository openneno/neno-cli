package main

import (
	"github.com/gobs/args"
	"github.com/peterh/liner"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"sort"
)

var (
	historyFile = "neno_command_history.txt"
)

func main() {

	app := &cli.App{
		Name:        "neno",
		Usage:       "在命令行中记录neno笔记",
		UsageText:   "neno [global options] command [command options] [arguments...]",
		Description: "neno 的一个命令行工具",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "token",
				Aliases: []string{"T"},
				Usage:   "Your github token",
				EnvVars: []string{"githubToken"},
			},
			&cli.StringFlag{
				Name:    "repo",
				Aliases: []string{"R"},
				Usage:   "Store your neno notes in a specific repo",
				EnvVars: []string{"githubRepo"},
			},
			&cli.StringFlag{
				Name:    "username",
				Aliases: []string{"U"},
				Usage:   "Your github username",
				EnvVars: []string{"githubUsername"},
			},
		},
		Commands: []*cli.Command{
			addCmd(),
			//ShowTags(),
			editCmd(),
			exitCmd(),
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				cli.ShowAppHelp(c)

				line := newLiner()
				defer closeLiner(line)

				for {
					if commandLine, err := line.Prompt("NENO > "); err == nil {
						line.AppendHistory(commandLine)

						cmdArgs := args.GetArgs(commandLine)
						if len(cmdArgs) == 0 {
							continue
						}

						s := []string{os.Args[0]}
						s = append(s, cmdArgs...)

						closeLiner(line)

						c.App.Run(s)

						line = newLiner()

					} else if err == liner.ErrPromptAborted || err == io.EOF {
						break
					} else {
						log.Print("Error reading line: ", err)
						continue
					}
				}
			}

			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
func newLiner() *liner.State {
	line := liner.NewLiner()

	line.SetCtrlCAborts(true)

	if f, err := os.Open(historyFile); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	return line
}

func closeLiner(line *liner.State) {
	if f, err := os.Create(historyFile); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
	line.Close()
}
