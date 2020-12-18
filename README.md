# go-panasonic

## Introduction
```go-panasonic``` is a Golang package and cli tool to control Panasonic AC equipment connected through the Panasonic Comfort Cloud (PCC). This library and cli tool has much of the same functionality as the Panasonic Comfort Cloud App.

Why? I got a bit bored with the App, it was too slow for me and I could not automate anything. So this package and cli tool solves that problem.

## CLI tool
You need a configuration file with the below content. The load the default file ```gopanasonic.yaml``` but it can also be passed with the cli flag ```-config [filepath]```.
```
username: [your PCC username]
password: [your PCC password]
device: [Panasonic device name, see -list command]
```

List all available Panasonic devices for account and add one of them to the configuration file.
```
$ go-panasonic -list
```

Some more examples
```
$ go-panasonic -status
$ go-panasonic -temp 19.5
$ go-panasonic -off
$ go-panasonic -on
$ go-panasonic -mode heat
$ go-panasonic -history week
```

```
$ go-panasonic -h
$ go-panasonic -version
```
