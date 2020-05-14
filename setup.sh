#!/bin/bash

#
# Setup script for setting up all requirements to run this code in this repository
# Author: Tarek Taha
#
# sudo apt install python3-pyside2.qtcore python3-pyside2.qtwidgets python3-pyside2.qtqml 
# sudo apt install python3-redis
sudo apt update && sudo apt-get upgrade
sudo apt-get install curl qml-module-qtcharts python3-pip golang python3-pyside2.qt3dcore python3-pyside2.qt3dinput python3-pyside2.qt3dlogic python3-pyside2.qt3drender python3-pyside2.qtcharts python3-pyside2.qtconcurrent python3-pyside2.qtcore python3-pyside2.qtgui python3-pyside2.qthelp python3-pyside2.qtlocation python3-pyside2.qtmultimedia python3-pyside2.qtmultimediawidgets python3-pyside2.qtnetwork python3-pyside2.qtopengl python3-pyside2.qtpositioning python3-pyside2.qtprintsupport python3-pyside2.qtqml python3-pyside2.qtquick python3-pyside2.qtquickwidgets python3-pyside2.qtscript python3-pyside2.qtscripttools python3-pyside2.qtsensors python3-pyside2.qtsql python3-pyside2.qtsvg python3-pyside2.qttest python3-pyside2.qttexttospeech python3-pyside2.qtuitools python3-pyside2.qtwebchannel python3-pyside2.qtwebsockets python3-pyside2.qtwidgets python3-pyside2.qtx11extras python3-pyside2.qtxml python3-pyside2.qtxmlpatterns python3-pyside2uic -y
python3 -m pip install pyyaml redis 
mkdir -p go/src/github.com/mzahmi
cd go/src/github.com/mzahmi 
git clone https://github.com/mzahmi/ventilator.git
cd ventilator
go get
cd
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo docker run --name some-redis -p 6379:6379 -d redis
sudo usermod -aG docker $USER
