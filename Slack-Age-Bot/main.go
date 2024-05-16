package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func printCommandEvents(analyticChannel <-chan *slacker.CommandEvent) {
	for event := range analyticChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Event)
		fmt.Println(event.Parameters)
		fmt.Println()
	}
}
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-7115043618226-7112185933221-1iOyCuO6sHLQzd6s3eoKf1yu")
	os.Setenv("SLACK_USER_TOKEN", "xapp-1-A073RM4GGG1-7117530227380-493d30a3dbee14ad8998437a43d4a5119289f0b120f33c25f3b1edac9f5b3568")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_USER_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("hello year of birth is <year>", &slacker.CommandDefinition{
		Description: "Yob calculator",
		Examples:    []string{"hello year of birth is 1990"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
				response.Reply("Invalid year")
				return
			}
			age := 2021 - yob
			response.Reply(fmt.Sprintf("Your age is %d", age))
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
