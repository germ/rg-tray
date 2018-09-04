RazerGenie-tray: A helper for RazerGenie configuration files

Installation:
Pacakges will be coming in the next bit, for now you must compile from source.
```
git clone https://github.com/germ/rg-tray.git	# Clone the repo
go get						# Download the dependancies
go build					# Build an executable
./rg-tray					# Launch!
```
All resources are bundled into the executable, simply launch the program and go.


Configuration:
Configs are stored in ~/.rg-tray.json, the following options are configurable and the defaults listed.

"GenieDir":"~/.local/share/razergenie/colours"
	- The directory to watch for matrix configurations

"Notray":false
	- Run in background without tray icon

"PollTime":0
	- Delay in seconds of how freqently to update keymap lighting. 0 disables


