// (c) 2017 - Bas Westerbaan <bas@westerbaan.name>
// You may redistribute this file under the conditions of the AGPLv3.

// twittermost is a mattermost bot that posts tweets of tweeps it follows
// on twitter.  See https://github.com/bwesterb/twittermost

package main

import (
	"github.com/nerzhul/twittermost/internal"
)

func main() {

	conf := internal.BotConf{}
	conf.Load()

	bot := internal.NewBot(conf)
	bot.Run()
	select {}
}
