package main

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	type Config struct {
		Twitter struct {
			ScreenName     string `json:"screen_name"`
			ConsumerKey    string `json:"consumer_key"`
			ConsumerSecret string `json:"consumer_secret"`
			AccessKey      string `json:"access_token"`
			AccessSecret   string `json:"access_secret"`
		} `json:"twitter"`
	}

	jsonFile, err := os.Open("./data/config.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	var cfg Config
	json.Unmarshal(jsonData, &cfg)

	if cfg.Twitter.ConsumerKey == "" || cfg.Twitter.ConsumerSecret == "" || cfg.Twitter.AccessKey == "" || cfg.Twitter.AccessSecret == "" {
		log.Fatal("Configuration missing")
	}

	oauth := oauth1.NewConfig(cfg.Twitter.ConsumerKey, cfg.Twitter.ConsumerSecret)
	token := oauth1.NewToken(cfg.Twitter.AccessKey, cfg.Twitter.AccessSecret)
	httpClient := oauth.Client(oauth1.NoContext, token)

	// main Twitter client
	client := twitter.NewClient(httpClient)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.User.ScreenName, "-> ", tweet.Text)
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Printf("%#v\n", dm)
	}
	// TODO: wire this up for likes i think?
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("%#v\n", event)
	}

	fmt.Println("Starting Stream...")

	// USER (quick test: auth'd user likes a tweet -> event)
	params := &twitter.StreamFilterParams{
		Track:         []string{cfg.Twitter.ScreenName},
		StallWarnings: twitter.Bool(true),
	}
	stream, _ := client.Streams.Filter(params)

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")

	stream.Stop()
}
