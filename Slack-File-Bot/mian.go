package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-7115043618226-7124424279686-jAVgFhN4e6WnvcosAyCYYZWM")
	os.Setenv("CHANNEL_ID", "C0741QHAQTA")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"ZIPL.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
