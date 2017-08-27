package main

import (
	"flag"
	"fmt"
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
	password = flag.String("password", "", "Password to connect to the server")
	channel  = flag.String("channel", "#goelo", "Channel to join")
)

func main() {
	flag.Parse()

	fmt.Printf("Connecting to server \033[1m%s\033[0m on port \033[1m%d\033[0m\n", *server, *port)

	irccon := irc.IRC("goelo", "goelo")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = false
	irccon.UseTLS = false
	irccon.Version = "Textual IRC Client: www.textualapp.com — v4.1.8 (Flavor: Pasilla de Oaxaca Chile) via ZNC 1.6.5+deb1 - http://znc.in"
	if *password != "" {
		irccon.Password = *password
	}

	// As soon as code 001 is received, join the channel
	irccon.AddCallback("001", func(e *irc.Event) {
		irccon.Join(*channel)
	})

	// If the bot is kicked from a channel, immediately rejoin it.
	// Maybe consider a timer or something for rejoins, or dont rejoin after X number of kicks.
	irccon.AddCallback("KICK", func(event *irc.Event) {
		go func(event *irc.Event) {
			irccon.Join(event.Arguments[0])
			fmt.Printf("I was kicked from %s by %s (%s)\n", event.Arguments[0], event.Nick, event.Arguments[2])
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
			fmt.Println("Added:", added)
		}(event)
	})

	err := irccon.Connect(fmt.Sprintf("%s:%d", *server, *port))
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}
