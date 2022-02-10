package main

import (
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func addCmd() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     "add note",
		UsageText: "添加一个笔记",
		Action: func(c *cli.Context) error {

			if c.NArg() <= 0 {
				return cli.ShowAppHelp(c)
			}
			return AddNeno(c, strings.Join(c.Args().Slice(), " "))
		},
	}
}
func AddNeno(c *cli.Context, content string) error {
	token := c.String("token")
	repo := c.String("repo")
	username := c.String("username")
	t := time.Now()
	createTime := t.Local().Format(time.RFC3339)
	createDate := t.Local().Format("2006-01-02")
	_id := NewObjectID().Hex()
	var tags []string
	tagsRex, _ := regexp.Compile(`#\S*`)
	tags = tagsRex.FindAllString(content, -1)
	for i, i2 := range tags {
		tags[i] = fmt.Sprintf(`"%s"`, i2)
	}
	tagString := strings.Join(tags, `,`)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s/%s.json", username, repo, createDate, _id)
	method := "PUT"
	neno := fmt.Sprintf(
		`{
        "content": "<p>%s</p>",
        "pureContent": "%s",
        "_id":"%s" ,
        "parentId": "",
        "source": "terminal",
        "tags": [%s],
        "images": [],
        "created_at": "%s",
        "sha": ""
	 }
`, content, content, _id, tagString, createTime)

	payload := strings.NewReader(fmt.Sprintf(`{
    "content": "%s",
    "message": "%s"
	}`, base64.StdEncoding.EncodeToString([]byte(neno)), fmt.Sprintf("[ADD] %s", content)))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

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

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == 401 {
		fmt.Println("TOKEN 错误")
		return err
	} else if res.StatusCode == 404 {
		fmt.Println("用户名或者仓库名错误")
	} else if res.StatusCode == 201 {
		fmt.Println("发送成功")
	} else {
		fmt.Println(string(body))
	}
	return nil
}
