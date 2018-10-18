package main

// import (
// 	"flag"
// 	"fmt"

// 	"github.com/davecgh/go-spew/spew"
// 	"github.com/marcusolsson/tui-go"
// 	"github.com/whyrusleeping/hellabot"
// 	"os"
// 	"time"
// )

// var serv = flag.String("server", "irc.oublinet.net:6697", "hostname and port for irc server to connect to")
// var nick = flag.String("nick", "elotenthree", "nickname for the bot")
// var chans = flag.String("channels", "#test", "Channels for the bot to join, separated by a comma")

// func main() {
// 	flag.Parse()

// 	sslOptions := func(bot *hbot.Bot) {
// 		bot.SSL = true
// 	}
// 	channels := func(bot *hbot.Bot) {
// 		bot.Channels = []string{*chans}
// 	}
// 	irc, err := hbot.NewBot(*serv, *nick, channels, sslOptions)
// 	if err != nil {
// 		panic(err)
// 	}

// 	history := tui.NewVBox()
// 	historyScroll := tui.NewScrollArea(history)
// 	historyScroll.SetAutoscrollToBottom(true)

// 	historyBox := tui.NewVBox(historyScroll)
// 	historyBox.SetBorder(true)

// 	input := tui.NewEntry()
// 	input.SetFocused(true)
// 	input.SetSizePolicy(tui.Expanding, tui.Maximum)

// 	inputBox := tui.NewHBox(input)
// 	inputBox.SetBorder(true)
// 	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

// 	chat := tui.NewVBox(historyBox, inputBox)
// 	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

// 	// TODO: Send the data to the irc server
// 	input.OnSubmit(func(e *tui.Entry) {
// 		irc.Send(fmt.Sprintf("PRIVMSG #test :%s", e.Text()))
// 		history.Append(tui.NewHBox(
// 			tui.NewLabel(time.Now().Format("15:04")),
// 			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("[%s]", *nick))),
// 			tui.NewLabel(e.Text()),
// 			tui.NewSpacer(),
// 		))
// 		input.SetText("")
// 	})

// 	root := tui.NewHBox(chat)

// 	ui, err := tui.New(root)
// 	if err != nil {
// 		panic(err)
// 	}

// 	ui.SetKeybinding("q", func() {
// 		ui.Quit()
// 		os.Exit(0)
// 	})

// 	ui.SetKeybinding("C-l", func() {
// 		ui.Quit()
// 		os.Exit(0)
// 	})

// 	// Listen for all irc messages and display them in the log panel
// 	// TODO: log them to the database
// 	irc.AddTrigger(hbot.Trigger{
// 		func(bot *hbot.Bot, m *hbot.Message) bool {
// 			return m.Command == "PRIVMSG"
// 		},
// 		func(irc *hbot.Bot, m *hbot.Message) bool {
// 			ui.Update(func() {
// 				history.Append(tui.NewHBox(
// 					tui.NewLabel(spew.Sprintf("%v", irc)),
// 					tui.NewLabel(spew.Sprintf("%v", m)),
// 					// tui.NewLabel(time.Now().Format("15:04")),
// 					// tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("[%s]", m.Name))),
// 					// tui.NewLabel(m.Content),
// 					// tui.NewSpacer(),
// 				))
// 			})
// 			return true
// 		},
// 	})

// 	go ui.Run()
// 	irc.Run()
// }
