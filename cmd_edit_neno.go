package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func editCmd() *cli.Command {
	return &cli.Command{
		Name:      "edit",
		Usage:     "edit",
		UsageText: "使用vim 进行编辑",
		Action: func(c *cli.Context) error {
			neno := execTextEditor()
			if neno == "panic]" {
				fmt.Println("请先安装vim")
				return nil
			}
			if neno == "" {
				fmt.Println("没有输入任何内容")
				return nil
			}
			return AddNeno(c, neno)

		},
	}
}
func execTextEditor() string {

	filePath := "" + time.Now().Format(time.RFC3339) + ".txt"
	cmd := exec.Command("vim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return "panic]"
	}

	tmpData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "panic]"
	}

	defer removeTmpFile(filePath, false)
	return string(tmpData)
}
func removeTmpFile(filePath string, removeAll bool) {
	realFilePath := filePath
	var err error
	err = os.Remove(realFilePath)
	if err != nil {
		panic(err)
	}
}
