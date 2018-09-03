// Razergenie-tray. A helper to deal with Razergenie configuration files
package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"gith
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
}
func onExit() {
}

