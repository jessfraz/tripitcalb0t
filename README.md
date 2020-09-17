# tripitcalb0t

[![make-all](https://github.com/jessfraz/tripitcalb0t/workflows/make%20all/badge.svg)](https://github.com/jessfraz/tripitcalb0t/actions?query=workflow%3A%22make+all%22)
[![make-image](https://github.com/jessfraz/tripitcalb0t/workflows/make%20image/badge.svg)](https://github.com/jessfraz/tripitcalb0t/actions?query=workflow%3A%22make+image%22)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/jessfraz/tripitcalb0t)
[![Github All Releases](https://img.shields.io/github/downloads/jessfraz/tripitcalb0t/total.svg?style=for-the-badge)](https://github.com/jessfraz/tripitcalb0t/releases)

Bot to automatically create Google Calendar events from TripIt flight data.

I have OCD about the layout of my calendar events so that is why I made my own bot.

**Table of Contents**

<!-- toc -->

- [Installation](#installation)
    + [Binaries](#binaries)
    + [Via Go](#via-go)
    + [Running with Docker](#running-with-docker)
- [Usage](#usage)
- [Setup](#setup)
  * [Google Calendar](#google-calendar)
  * [TripIt](#tripit)

<!-- tocstop -->

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
    -v /path/to/.tripitcalb0t/google.json:/root/.tripitcalb0t/google.json:ro \
    --name tripitcalb0t \
    -e "TRIPIT_USERNAME=your_username" \
    -e "TRIPIT_PASSWORD=59f6asdfasdfasdf0" \
    -e "GOOGLE_CALENDAR_ID=your_google_calendar_id" \
    r.j3ss.co/tripitcalb0t --interval 1m
```

## Usage

```console
$ tripitcalb0t -h
tripitcalb0t -  Bot to automatically create Google Calendar events from TripIt flight data.

Usage: tripitcalb0t <command>

Flags:

  --calendar         Calendar name to add events to (or env var GOOGLE_CALENDAR_ID)
  -d                 Enable debug logging (default: false)
  --google-keyfile   Path to Google Calendar keyfile (default: ~/.tripitcalb0t/google.json)
  --interval         Update interval (ex. 5ms, 10s, 1m, 3h) (default: 1m0s)
  --once             Run once and exit, do not run as a daemon (default: false)
  --past             Include past trips (default: false)
  --tripit-password  TripIt Password for authentication (or env var TRIPIT_PASSWORD)
  --tripit-username  TripIt Username for authentication (or env var TRIPIT_USERNAME)

Commands:

  version  Show the version information.
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
