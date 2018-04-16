# tripitcalb0t

[![Travis CI](https://travis-ci.org/jessfraz/tripitcalb0t.svg?branch=master)](https://travis-ci.org/jessfraz/tripitcalb0t)

Bot to automatically create Google Calendar events from TripIt flight data.

I have OCD about the layout of my calendar events so that is why I made my own
bot.

## Installation

#### Binaries

- **darwin** [386](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-darwin-386) / [amd64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-darwin-amd64)
- **freebsd** [386](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-freebsd-386) / [amd64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-freebsd-amd64)
- **linux** [386](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-linux-386) / [amd64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-linux-amd64) / [arm](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-linux-arm) / [arm64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-linux-arm64)
- **solaris** [amd64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-solaris-amd64)
- **windows** [386](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-windows-386) / [amd64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.0.0/tripitcalb0t-windows-amd64)


#### Via Go

```bash
$ go get github.com/jessfraz/tripitcalb0t
```

#### Running with Docker

```console
$ docker run --restart always -d \
    -v /etc/localtime:/etc/localtime:ro \
    --name tripitcalb0t \
    -e "TRIPIT_USERNAME=your_username" \
    -e "TRIPIT_TOKEN=59f6asdfasdfasdf0" \
    r.j3ss.co/tripitcalb0t --interval 1m
```

## Usage

```console
 _        _       _ _            _ _      ___  _
| |_ _ __(_)_ __ (_) |_ ___ __ _| | |__  / _ \| |_
| __| '__| | '_ \| | __/ __/ _` | | '_ \| | | | __|
| |_| |  | | |_) | | || (_| (_| | | |_) | |_| | |_
 \__|_|  |_| .__/|_|\__\___\__,_|_|_.__/ \___/ \__|
           |_|

 Bot to automatically create Google Calendar events from TripIt flight data.
 Version: v0.0.0
 Build: 1f12d9f

  -d    run in debug mode
  -interval string
        update interval (ex. 5ms, 10s, 1m, 3h) (default "1m")
  -once
        run once and exit, do not run as a daemon
  -tripit-token string
        TripIt Token for authentication (or env var TRIPIT_TOKEN)
  -tripit-username string
        TripIt Username for authentication (or env var TRIPIT_USERNAME)
  -v    print version and exit (shorthand)
  -version
        print version and exit
```
