# Go Hue #

Simple CLI script written in Go to control your Philips Hue lights.

## Usage ##

```
Golang CLI application to communicate with your Philips Hue light system

Check https://github.com/jordyv/hue for more information

Usage:
  hue [command]

Available Commands:
  help        Help about any command
  lights      Actions for your lights

Flags:
      --config string   config file (default is $HOME/.hue.yaml)
  -h, --help            help for hue
      --ip string       IP address of your Hue bridge
      --token string    Token of your Hue bridge
  -v, --verbose         Verbose logging

Use "hue [command] --help" for more information about a command.
```

```
# Example: 
$ hue lights on
$ hue lights off
```

## Installation ##

```
go get https://github.com/jordyv/hue
```

## Configuration ##

Create a `.hue.yaml` file in your home directory. Check `hue.yaml.example` for an example.

```yaml
ip: 192.123.123.123
token: <your user token for the Hue API>
sceneID: <ID of the scene you want to turn on and off>
```

### Create an Hue API user token ###

```
curl -X POST --data '{ "devicetype": "go_hue#api" }' http://<IP of your Hue bridge>/api
```
You can find the API token as `username` in the response.
