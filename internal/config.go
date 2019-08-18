package internal

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

type BotConf struct {
	Url      string // URL to mattermost instance
	DataPath string // path to data

	// mattermost settings
	User         string
	Email        string
	Password     string
	Token        string
	Team         string
	Channel      string
	DebugChannel string

	// twitter settings
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
	MaxTweets      int
	CheckInterval  int

	ServicePort         int
	ServiceAllowedToken string
}

func (c *BotConf) loadDefaultConfiguration() {
	c.Url = "http://localhost:8065"
	c.DataPath = "mattermost.json"

	c.User = "bot"
	c.Password = "secret"
	c.Token = ""
	c.Team = "team"
	c.Channel = "twitter"
	c.DebugChannel = "twitter-debug"

	c.ConsumerKey = ""
	c.ConsumerSecret = ""
	c.AccessToken = ""
	c.AccessSecret = ""
	c.MaxTweets = 20
	c.CheckInterval = 120

	c.ServicePort = 8080
	c.ServiceAllowedToken = ""
}

func (c *BotConf) loadFromFile() error {
	var confPath = "conf/config.json"

	//Check whether old config file exists
	if _, err := os.Stat("config.json"); !os.IsNotExist(err) {
		// fall back to old config file location
		confPath = "config.json"
	}

	// Parse cmdline flags
	flag.StringVar(&confPath, "config", confPath,
		"Path to configuration file")
	flag.Parse()

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		log.Println("Configuration file not found, ignoring.")
		return nil
	}
	// Read config file
	buf, err := ioutil.ReadFile(confPath)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not read config file %s: %s", confPath, err))
	}

	if err := json.Unmarshal(buf, c); err != nil {
		return errors.New(fmt.Sprintf("Could not parse config file: %s", err))
	}

	log.Println("Configuration file loaded.")
	return nil
}

func (c *BotConf) loadFromEnv() {
	err := envconfig.Process("", c)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Loaded configuration from env vars")
}

func (c *BotConf) Load() {
	c.loadDefaultConfiguration()
	err := c.loadFromFile()
	if err != nil {
		log.Printf("Failed to read config from file: %s\n", err)
	}

	c.loadFromEnv()

	mattermostKubeUrl := os.Getenv("MATTERMOST_PORT")
	if len(mattermostKubeUrl) > 0 {
		log.Printf("[k8s discovery] MATTERMOST_PORT url found: %s\n", mattermostKubeUrl)
		mattermostParsedURL, err := url.Parse(mattermostKubeUrl)
		if err != nil {
			log.Fatalf("Unable to parse MATTERMOST_PORT URL: %v\n", err)
		}

		if mattermostParsedURL.Scheme == "tcp" {
			mattermostParsedURL.Scheme = "http"
		}

		c.Url = fmt.Sprintf("%s://%s", mattermostParsedURL.Scheme, mattermostParsedURL.Host)
		log.Printf("Mattermost URL set to: %s\n", c.Url)
	}
}
