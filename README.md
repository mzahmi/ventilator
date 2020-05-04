# Ventinator software

## Install

## Pre-requisites

On Ubuntu / Debian

### Install GoLang

```
sudo apt install golang
```

### Install QML GoLang Wrapper
```
export GO111MODULE=off; go get -v github.com/therecipe/qt/cmd/... && $(go env GOPATH)/bin/qtsetup test && $(go env GOPATH)/bin/qtsetup -test=false
echo export PATH=$PATH:/home/$USER/go/bin/ >> ~/.bashrc
```
### Install GoLang Dependencies

```
go get github.com/tarm/serial
```

## Install Ventilator Software 

```
git clone https://github.com/mzahmi/ventilator/
```

All code and details related to the open ventilator project.

```
$ cd ventilator/gui
qtdeploy test desktop .

```
