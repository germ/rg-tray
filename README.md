### rg-tray 

A helper for RazerGenie configuration files and colorschemes
for Razer keyboards that lives in your tray.

##### Screenshot:

![It aint much right now :)](https://i.imgur.com/vrRmwSz.png)

##### Installation:

Pacakges will be coming in the next bit, for now you must compile from source.
```
git clone https://github.com/germ/rg-tray.git	# Clone the repo
go get						# Download the dependancies
go build					# Build an executable
./rg-tray					# Launch! (or move into $PATH and add to autostart)
```
All resources are bundled into the executable, simply launch the program and go.

##### Color Schemes:

rg-tray uses schemes exported from [RazerGenie](https://github.com/z3ntu/RazerGenie)
the default configuration uses 

##### Configuration:

Configuration is stored in ~/.rg-tray.json, the following options are configurable
If no config is present, these defaults are used. Notray and Polltime are not 
implemented yet
```
{
  "GenieDir":"~/.local/share/razergenie/colours", # The directory to watch for matrix configurations
  "Notray":false,                                 # Run in background without tray icon
  "PollTime":0,                                   # seconds between config reloads, 0 to disable
  "AltIcon":false,                                # use inverted icon
}
```

