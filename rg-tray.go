// Razergenie-tray. A helper to deal with Razergenie configuration files
package main

import (
	"fmt"
	"os"

	razer "github.com/germ/go-razer"
	"github.com/getlantern/systray"
)

func main() {
	fmt.Println(Config)
	c, err := readProfile(Config.Profile)
	fmt.Println(c, err)
	os.Exit(1)

	fmt.Println("Probing Devices")
	devices, _ := razer.Devices()
	for _, v := range devices {
		fmt.Println(v, v.String(), v.FullName(), v.Type())
	}

	systray.Run(trayReady, trayExit)
}

func trayReady() {
	systray.SetIcon(icon)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	_ = mQuit
}

func trayExit() {
}
