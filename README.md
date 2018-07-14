# tripitcalb0t

[![Travis CI](https://travis-ci.org/jessfraz/tripitcalb0t.svg?branch=master)](https://travis-ci.org/jessfraz/tripitcalb0t)

Bot to automatically create Google Calendar events from TripIt flight data.

I have OCD about the layout of my calendar events so that is why I made my own
bot.

## Installation

#### Binaries

- **darwin** [386](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-darwin-386) / [amd64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-darwin-amd64)
- **freebsd** [386](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-freebsd-386) / [amd64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-freebsd-amd64)
- **linux** [386](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-linux-386) / [amd64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-linux-amd64) / [arm](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-linux-arm) / [arm64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-linux-arm64)
- **solaris** [amd64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-solaris-amd64)
- **windows** [386](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-windows-386) / [amd64](https://github.com/jessfraz/tripitcalb0t/releases/download/v0.1.2/tripitcalb0t-windows-amd64)


#### Via Go

```bash
$ go get github.com/jessfraz/tripitcalb0t
```

#### Running with Docker

```console
$ docker run --restart always -d \
    -v /etc/localtime:/etc/localtime:ro \
    -v /path/to/.tripitcalb0t/google.json:/.tripitcalb0t/google.json:ro \
    --name tripitcalb0t \
    -e "TRIPIT_USERNAME=your_username" \
    -e "TRIPIT_TOKEN=59f6asdfasdfasdf0" \
    -e "GOOGLE_CALENDAR_ID=your_google_calendar_id" \
    r.j3ss.co/tripitcalb0t --interval 1m
```

## Usage

```console
$ tripitcalb0t -h
 _        _       _ _            _ _      ___  _
| |_ _ __(_)_ __ (_) |_ ___ __ _| | |__  / _ \| |_
| __| '__| | '_ \| | __/ __/ _` | | '_ \| | | | __|
| |_| |  | | |_) | | || (_| (_| | | |_) | |_| | |_
 \__|_|  |_| .__/|_|\__\___\__,_|_|_.__/ \___/ \__|
           |_|

 Bot to automatically create Google Calendar events from TripIt flight data.
 Version: v0.1.2
 Build: cb952a5

  -calendar string
        Calendar name to add events to (or env var GOOGLE_CALENDAR_ID)
  -d    run in debug mode
  -google-keyfile string
        Path to Google Calendar keyfile (default "~/.tripitcalb0t/google.json")
  -interval duration
        update interval (ex. 5ms, 10s, 1m, 3h) (default 1m0s)
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

## Setup

### Google Calendar

1. Enable the API: To get started using Calendar API v3, you need to 
    first create a project in the 
    [Google API Console](https://console.developers.google.com),
    enable the API, and create credentials.

    Follow the instructions 
    [for step enabling the API here](https://developers.google.com/calendar/quickstart/go).

2. Add the new service account to the Google Calendar account with 
    [Read & Write](https://support.google.com/analytics/answer/2884495) 
    permission.

    The newly created service account will have an email address that looks
    similar to: `quickstart@PROJECT-ID.iam.gserviceaccount.com`.

    Use this email address to 
    [add a user](https://support.google.com/analytics/answer/1009702) to the 
    Google Calendar view you want to access via the API. 

### TripIt

To use this, you must enable "Web Authentication" on your account. You can
follow the steps to do that 
[here](https://tripit.github.io/api/doc/v1/#authentication_section).
