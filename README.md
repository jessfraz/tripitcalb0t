# tripitcalb0t

[![Travis CI](https://img.shields.io/travis/jessfraz/tripitcalb0t.svg?style=for-the-badge)](https://travis-ci.org/jessfraz/tripitcalb0t)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/jessfraz/tripitcalb0t)
[![Github All Releases](https://img.shields.io/github/downloads/jessfraz/tripitcalb0t/total.svg?style=for-the-badge)](https://github.com/jessfraz/tripitcalb0t/releases)

Bot to automatically create Google Calendar events from TripIt flight data.

I have OCD about the layout of my calendar events so that is why I made my own bot.

 * [Installation](README.md#installation)
      * [Binaries](README.md#binaries)
      * [Via Go](README.md#via-go)
      * [Running with Docker](README.md#running-with-docker)
 * [Usage](README.md#usage)
 * [Setup](README.md#setup)
   * [Google Calendar](README.md#google-calendar)
   * [TripIt](README.md#tripit)

## Installation

#### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/jessfraz/tripitcalb0t/releases).

#### Via Go

```console
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
 Version: v0.1.3
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
