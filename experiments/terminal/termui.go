package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	ui "github.com/gizak/termui"
)

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	ls := ui.NewList()
	ls.Border = true
	ls.BorderLabel = "Logs"
	ls.Items = []string{
		"[19:19:25] ( nick) jodell: didnt do the breathatarian diet",
		"[22:43:05] ( gabriel) holy fuckin lol https://twitter.com/cocksailor/status/1051588994499133441",
		"[23:26:00] ( moonpolysoft) man fuckin kenji lopez-alt retweeted me today and his followers are a tire fire",
	}
	ls.Height = ui.TermHeight() - 3

	par := ui.NewPar(" > ")
	par.Height = 3
	par.BorderLabel = "Input"

	// build layout
	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, ls)),
		ui.NewRow(ui.NewCol(12, 0, par)),
	)

	// calculate layout
	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Handle("q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd", func(e ui.Event) {
		fmt.Println(spew.Printf("%v", e))
		ui.Render(ui.Body)
	})

	ui.Handle("<Resize>", func(e ui.Event) {
		payload := e.Payload.(ui.Resize)
		ui.Body.Width = payload.Width
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Loop()
}
