package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"os/user"
	"strings"

	razer "github.com/germ/go-razer"
	colours "github.com/go-playground/colors"
)

const configPath = "~/.rgtray.json"

var Config Conf

type Conf struct {
	GenieDir string
	PollTime int
	Notray   bool
	Profile  string
}

var defaultConf = Conf{
	GenieDir: "~/.local/share/razergenie/colours/",
	PollTime: 120,
	Notray:   false,
	Profile:  "default",
}

type Colours struct {
	Name   string
	Author string
	Dev    map[string]Device
}

type Device struct {
	DevType string `json:"Type"`
	Matrix  [][]color.Color

	// Other types device properties to be added as implemented
	// DevType will act as switch
	// All types: keyboard, mouse, core, mug, keypad, headset, mousemat
}

func init() {
	// Try loading conf, use default if not found
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		Config = defaultConf
	} else {
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(data, &Config)
		if err != nil {
			panic(err)
		}
	}
}

// Open/parse a JSON scheme and apply to devices
func readProfile(profileName string) (c Colours, err error) {
	// Sustitute home dir, read in profile
	u, _ := user.Current()
	path := Config.GenieDir + profileName + ".json"
	path = strings.Replace(path, "~", u.HomeDir, -1)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &c)
	return
}

func (c *Colours) apply() (err error) {
	// Get list of devices, check for matches
	devices, err := razer.Devices()
	if err != nil {
		return
	}

	for _, v := range devices {
		if dev, ok := c.Dev[v.FullName()]; ok {
			// Found device with configuration
			switch dev.DevType {
			case "keyboard":
				var keys []razer.KeySet

				// Generate stuff to send
				for row, v := range dev.Matrix {
					var curRow razer.KeySet
					for key, rgb := range v {
						curRow = append(curRow, razer.Key{
							Row:   row,
							Col:   key,
							Color: rgb,
						})
					}
					keys = append(keys, curRow)
				}

				//Send it!
				for 
			default:
				fmt.Println("Your device is not supported at the moment")
				fmt.Println("Open a issue and I'll get to it :)")
			}
			// TODO: Add in other types of devices
		} else {
			fmt.Println("Device not configured: ", v.FullName)
		}
	}
	return
}

// Custom decoder, needed because of custom name :(
// And yeah, it's super fucking gross.
// Normally we just define a struct and encoding/json takes
// care of the type infrence.
func (c *Colours) UnmarshalJSON(data []byte) (err error) {
	c.Dev = make(map[string]Device)
	var raw map[string]interface{}

	// Decode json
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return
	}

	// Pull and delete metainfo
	if v, ok := raw["Author"]; ok {
		c.Author = v.(string)
		delete(raw, "Author")
	}

	if v, ok := raw["Name"]; ok {
		c.Name = v.(string)
		delete(raw, "Name")
	}

	// Loop over devices
	for devName, devInfo := range raw {
		rawDev := devInfo.(map[string]interface{})
		var NewDevice Device

		if v, ok := rawDev["Type"]; ok {
			NewDevice.DevType = v.(string)
			delete(rawDev, "Type")
		}

		// Parse device sensitive stuff
		// Only adding valid devices
		switch NewDevice.DevType {
		case "keyboard":
			// Loop over the matrix and fill
			if v, ok := rawDev["Matrix"]; ok {
				matRaw := v.([]interface{})

				for _, row := range matRaw {
					rowRaw := row.([]interface{})
					var sRow []color.Color
					for _, key := range rowRaw {
						keyColour, err := colours.ParseHEX(key.(string))
						if err != nil {
							fmt.Println("Error parsing key: ", key.(string))
							continue
						}

						rgb := keyColour.ToRGBA()
						sRow = append(sRow, color.RGBA{R: rgb.R, G: rgb.G, B: rgb.B, A: 128})
					}
					NewDevice.Matrix = append(NewDevice.Matrix, sRow)
				}
				delete(rawDev, "Matrix")
			}

		default:
			fmt.Println("Unrecognized Device Type", devName)
			fmt.Println("Please open a ticket and I'll try and support it :)")
		}

		if len(rawDev) != 0 {
			fmt.Println("Unhandled elements in device", devName)
			fmt.Println(rawDev)
		}

		c.Dev[devName] = NewDevice
	}

	if len(raw) != 0 {
		fmt.Println("Unhandled elements in colorMap")
		fmt.Println(raw)
	}

	return
}
