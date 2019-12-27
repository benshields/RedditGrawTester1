package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type reminderBot struct {
	bot reddit.Bot
}

func (r *reminderBot) Post(p *reddit.Post) error {
	if strings.Contains(p.Author, "grawbot1") && strings.Contains(p.SelfText, "remind me of this post") {
		fmt.Printf("Found a matching post titled [%s] by [%s]\n", p.Title, p.Author)
		<-time.After(10 * time.Second)
		return r.bot.SendMessage(
			p.Author,
			fmt.Sprintf("Reminder: %s", p.Title),
			"You've been reminded!",
		)
	}
	return nil
}

func main() {
	if bot, err := reddit.NewBotFromAgentFile("grawbot1.agent", 0); err != nil {
		fmt.Println("Failed to create bot handle: ", err)
	} else {
		cfg := graw.Config{Subreddits: []string{"bottesting"}}
		handler := &reminderBot{bot: bot}
		if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
			fmt.Println("Failed to start graw run: ", err)
		} else {
			fmt.Println("graw run failed: ", wait())
		}
	}
}
