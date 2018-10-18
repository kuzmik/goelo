package main

import (
	"flag"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/getsentry/raven-go"
	"github.com/golang/glog"
	"github.com/thoj/go-ircevent"
)

/*
&irc.Event{
Code:"PRIVMSG",
Raw:":nick!nick@moron5.com PRIVMSG #test :imitate",
Nick:"nick",
Host:"moron5.com",
Source:"nick!nick@moron5.com",
User:"nick",
Arguments:[]string{"#test", "imitate"},
Connection:(*irc.Connection)(0xc4200a4000)
}
*/

var (
	server   = flag.String("server", "irc.slashnet.org", "IRC server to connect to")
	port     = flag.Int("port", 6667, "Port to connect to")
	ssl      = flag.Bool("ssl", false, "Connect with SSL")
	password = flag.String("password", "", "Password to connect to the server")
	channels = flag.String("channels", "#test", "Channels to join")
)

func main() {
	raven.SetDSN("https://d7ad690e9623491cb19250dbe64fc401@sentry.io/1213124")

	flag.Parse()

	fmt.Printf("Connecting to server \033[1m%s\033[0m on port \033[1m%d\033[0m\n", *server, *port)

	irccon := irc.IRC("goelo", "goelo")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = false
	irccon.UseTLS = *ssl
	irccon.Version = "Textual IRC Client: www.textualapp.com — v4.1.8 (Flavor: Pasilla de Oaxaca Chile) via ZNC 1.6.5+deb1 - http://znc.in"
	if *password != "" {
		irccon.Password = *password
	}

	// As soon as code 001 is received, join the channel
	irccon.AddCallback("001", func(e *irc.Event) {
		irccon.Join(*channels)
	})

	// If the bot is kicked from a channel, immediately rejoin it.
	// Maybe consider a timer or something for rejoins, or dont rejoin after X number of kicks.
	irccon.AddCallback("KICK", func(event *irc.Event) {
		go func(event *irc.Event) {
			irccon.Join(event.Arguments[0])
			glog.Infof("I was kicked from %s by %s (%s)\n", event.Arguments[0], event.Nick, event.Arguments[2])
		}(event)
	})

	// Handle incoming message
	irccon.AddCallback("PRIVMSG", func(event *irc.Event) {
		b := Bebot{}

		// "talk about $nick"
		if event.Message() == "imitate" {
			message := b.GetLog(event.Nick)
			irccon.Privmsg(event.Arguments[0], message)
		}

		// Log all messages
		go func(event *irc.Event) {
			added := b.AddLog(*server, event.Arguments[0], event.Nick, event.Message())
			glog.Infof("Added new log entry: ", added)
		}(event)
	})

	err := irccon.Connect(fmt.Sprintf("%s:%d", *server, *port))
	if err != nil {
		glog.Fatal(err)
		raven.CaptureErrorAndWait(err, nil)
		return
	}
	irccon.Loop()
}
