# OpenVentAid Ventilator [![Build Status](https://travis-ci.org/mzahmi/ventilator.svg?branch=master)](https://travis-ci.org/mzahmi/ventilator)

## Install

## Pre-requisites

On Ubuntu / Debian

### Install GoLang

```
sudo apt install golang
```

### Install GoLang Dependencies

```
go get github.com/tarm/serial
```

### Install required build dependencies

You will need to install plenty of packages to be able to build Qt. Some of the Qt features are optional, for example support for various databases and if you don't need a specific feature you can skip building the support. Or the other way around, if you need a specific feature you might need to install more packages. See the table below for a list of some optional features and the required development packages you need to install. But first, start by updating your package cache so everything is fresh:

```
sudo apt-get update

sudo apt-get install build-essential libfontconfig1-dev libdbus-1-dev libfreetype6-dev libicu-dev libinput-dev libxkbcommon-dev libsqlite3-dev libssl-dev libpng-dev libjpeg-dev libglib2.0-dev libraspberrypi-dev
```

## Install github.com/mzahmi/ventilatorilator Software 

```
git clone https://github.com/mzahmi/ventilator/
git submodule init
git submodule update
```


