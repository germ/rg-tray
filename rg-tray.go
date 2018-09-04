// Razergenie-tray. A helper to deal with Razergenie configuration files
package main

import (
	"fmt"

	"github.com/getlantern/systray"
)

const configPath = "~/.rgtray.json"

var Config Conf
var Scheme Colours

func main() {
	// Aight, so this is a bit weird. We call Run
	// which starts the GUI and fires trayReady. Eventually
	// The user clicks quit and trayExit fires
	// Until then this blocks.
	// Init is in config.go
	systray.Run(trayReady, trayExit)
}

func trayReady() {
	systray.SetIcon(icon)
	systray.SetTitle("rg-tray configurator")
	systray.SetTooltip("rg-tray")
	mDisable := systray.AddMenuItem("Disable", "Turn off all LEDs")
	mEnable := systray.AddMenuItem("Enable", "Turn on all LEDs")
	mQuit := systray.AddMenuItem("Quit", "Quit")

	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				fmt.Println("Quitting...")
				systray.Quit()
			case <-mDisable.ClickedCh:
				fmt.Println("Disabling...")
				mDisable.Disable()
				mEnable.Enable()
				changeAllBrightness(0.0)
			case <-mEnable.ClickedCh:
				fmt.Println("Enabling...")
				mDisable.Enable()
				mEnable.Disable()
				changeAllBrightness(100.0)
			}
		}
	}()
}

func trayExit() {
	systray.Quit()
}
