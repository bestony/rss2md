package main

import (
	"fmt"
	"os"

	"github.com/ungerik/go-rss"
)

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	channel, err := rss.Read(args[1])
	if err != nil {
		fmt.Println(err)
	}
	var content string
	content = channel.Title + "[Link](" + channel.Link + ")\n======\n"
	for _, item := range channel.Item {
		var link string
		link = "- [" + item.Title + "](" + item.Link + ")\n"
		content = content + link
	}
	dstFile, err := os.Create(channel.Title + ".md")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(content + "\n")
	fmt.Println("convert success！")
}
