// (c) 2017 - Bas Westerbaan <bas@westerbaan.name>
// You may redistribute this file under the conditions of the AGPLv3.

// twittermost is a mattermost bot that posts tweets of tweeps it follows
// on twitter.  See https://github.com/bwesterb/twittermost

package main

import (
	"encoding/json"
	"flag"
	"github.com/nerzhul/twittermost/internal"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var confPath = "conf/config.json"
	conf := internal.BotConf{
		DataPath:      "mattermost.json",
		Channel:       "town-square",
		DebugChannel:  "test",
		MaxTweets:     20,
		CheckInterval: 120,
	}

	//Check whether old config file exists
	if _, err := os.Stat("config.json"); !os.IsNotExist(err) {
		// fall back to old config file location
		confPath = "config.json"
	}

	// Parse cmdline flags
	flag.StringVar(&confPath, "config", confPath,
		"Path to configuration file")
	flag.Parse()

	// Read config file
	buf, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatalf("Could not read config file %s: %s", confPath, err)
	}

	if err := json.Unmarshal(buf, &conf); err != nil {
		log.Fatalf("Could not parse config file: %s", err)
	}

	bot := internal.NewBot(conf)
	bot.Run()
	select {}
}
