package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/TheL1ne/Songrequests/UICreator"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.SetContent(container.NewVBox(UICreator.MakeUI()))
	w.ShowAndRun()
}
