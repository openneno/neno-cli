package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
)

func ShowTags() *cli.Command {
	return &cli.Command{
		Name:      "tags",
		Usage:     "tags tagName",
		UsageText: "查询一个tag或者不加参数时,显示所有的tag",
		Action: func(c *cli.Context) error {

			if c.NArg() <= 0 {
				err := cli.ShowAppHelp(c)
				if err != nil {
					return err
				}
			}
			return getTags(c, c.Args().First())
		},
	}
}
func getTags(c *cli.Context, tag string) error {
	token := c.String("token")
	repo := c.String("repo")
	username := c.String("username")

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/.json", username, repo)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("authorization", fmt.Sprintf("token %s", token))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
