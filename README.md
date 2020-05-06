# github.com/mzahmi/ventilatorinator software [![Build Status](https://travis-ci.org/mzahmi/github.com/mzahmi/ventilatorilator.svg?branch=master)](https://travis-ci.org/mzahmi/github.com/mzahmi/ventilatorilator)

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

## Install github.com/mzahmi/ventilatorilator Software 

```
git clone https://github.com/mzahmi/ventilator/
```

All code and details related to the open github.com/mzahmi/ventilatorilator project.

```
$ cd github.com/mzahmi/ventilatorilator/gui
qtdeploy test desktop .

```

# Build Qt (with Extra Modules) on a Raspberry-Pi 4

Tutorial from <https://www.tal.org/tutorials/building-qt-512-raspberry-pi> but I found some bugs so here is an updated version. This will build qt with the necissary dependancies to run the github.com/mzahmi/ventilator software on a raspberry pi 4 running raspbien.

### Download the Qt 5.12.7 source archive

Download the single source tar file from download.qt.io, [version 5.12.7](http://download.qt.io/official_releases/qt/5.12/5.12.7/single/qt-everywhere-src-5.12.7.tar.xz). The package is quite large, **482MB**, so depending on your bandwidth it might take some time to download, grab a coffe or continue with the dependecies while the Qt source archive is downloading.

```
wget http://download.qt.io/official_releases/qt/5.12/5.12.7/single/qt-everywhere-src-5.12.7.tar.xz
```

### Check archive MD5 hash

Just to make sure the download was successful and is what it should be, check the [MD5 hash of the archive](http://download.qt.io/official_releases/qt/5.12/5.12.7/single/md5sums.txt) with:

```
md5sum qt-everywhere-src-5.12.7.tar.xz
```

It should be:

```
ce2c5661c028b9de6183245982d7c120  qt-everywhere-src-5.12.7.tar.xz
```

### Un-tar the source archive

Un-tar the source archive in a suitable location, with enough free space (**\~2.8GB)**. This will take around *7-13 minutes* on a Raspberry Pi 3+, depending on SD card speed, and even more on an older model, so go grab of coffe again. You can of course start installation of the build dependencies while you wait.

```
tar xf qt-everywhere-src-5.12.7.tar.xz
```

### Setup Qt mkspecs configuration files

Qt build is configured trough the configure script, but platform and device specifc settings are set in mkspecs configuration files. Qt includes mkspecs for the Raspberry Pi but they are unfortunately setup for cross-compilation environments only and can not be used for native building without editing or the need for faking a cross compilation environment and that has its own issues.

Fortunately suitable mkspecs files for the various Pi versions [are available in our github](https://github.com/oniongarlic/qt-raspberrypi-configuration.git), so the next step is to clone that repository and install the mkspecs files into the Qt source tree so that they can be used.

Clone the configuration repository:

```
git clone https://github.com/oniongarlic/qt-raspberrypi-configuration.git
```

change into the cloned repository and run (adjust DESTDIR in case your Qt sources are extracted somewhere else):

```
cd qt-raspberrypi-configuration && make install DESTDIR=../qt-everywhere-src-5.12.7
```

### Install required build dependencies

You will need to install plenty of packages to be able to build Qt. Some of the Qt features are optional, for example support for various databases and if you don't need a specific feature you can skip building the support. Or the other way around, if you need a specific feature you might need to install more packages. See the table below for a list of some optional features and the required development packages you need to install. But first, start by updating your package cache so everything is fresh:

```
sudo apt-get update

sudo apt-get install build-essential libfontconfig1-dev libdbus-1-dev libfreetype6-dev libicu-dev libinput-dev libxkbcommon-dev libsqlite3-dev libssl-dev libpng-dev libjpeg-dev libglib2.0-dev libraspberrypi-dev
```

### Create a shadow build directory outside of the Qt source tree

We will build Qt outside of the source tree, this way you can easily have different build version and easily also start over in case of any issues. You build location can be anywhere where there is enough space, for example an USB stick in case you are running out on your SD card. Remember to adjust any paths in the commands.

```
mkdir build
cd build
```

### Configure the Qt build environment

We configure Qt for a native build and set device specific settings, like libraries, default [QPA](http://doc.qt.io/qt-5/qpa.html) plugin to use, etc directly on the configure command line. This makes it also easy to optimize for any specific Raspberry Pi board type, see table for suitable CFLAGS & CXXFLAGS for your target board. Configure run takes \~10 minutes as it will first compile qmake and then run various tests and feature checks.

*before compiling* edit the following file

```
.../qt-everywhere-src-5.12.7/qtscript/src/3rdparty/javascriptcore/JavaScriptCore/wtf/Platform.h
```

by adding the following lines (without the +)

```
...
#elif defined(__ARM_ARCH_7A__) \
    || defined(__ARM_ARCH_7R__)
#define WTF_ARM_ARCH_VERSION 7
                                                                                                                                                                                           +#elif defined(__ARM_ARCH_8A__)
+#define WTF_ARM_ARCH_VERSION 8

/* RVCT sets _TARGET_ARCH_ARM */
...
```

now run

```
PKG_CONFIG_LIBDIR=/usr/lib/arm-linux-gnueabihf/pkgconfig:/usr/share/pkgconfig \
../qt-everywhere-src-5.12.7/configure -platform linux-rpi-g++ \
-v \
-opengl es2 -eglfs \
-no-gtk \
-opensource -confirm-license -release \
-reduce-exports \
-force-pkg-config \
-nomake examples -no-compile-examples \
-skip qtwayland \
-skip qtwebengine \
-no-feature-geoservices_mapboxgl \
-qt-pcre \
-no-pch \
-ssl \
-evdev \
-system-freetype \
-fontconfig \
-glib \
-prefix /opt/Qt5.12 \
-qpa eglfs \
-platform linux-rpi4-v3d-g++
```

### Raspberry Pi platform options

Select platform according to the Pi model you are going to run Qt on. You can build on a fast Pi 4/3+ and the move the files over to a slower version if needed. Note that for a Raspberry Pi 4 you must use KMS, Broadcom EGL is not supported.

### Check the configure result

When the Qt configuration scripts has finnished succesfully it will print out a configuration summary and store the result in the *config.summary* file, also all the raw checks will be in the *config.log* file. Check that all options and features you needed are found and enabled in the summary output. If not, check the *config.log* file for a reason.

#### Broadcom EGLFS

Make sure that the configure script detects Raspberry Pi EGLFS when building for Raspberry Pi 1,2 or 3. Look for the following output or check the log file *config.summary*:

```
  EGLFS .................................. yes
  EGLFS details:
...
    EGLFS Rasberry Pi .................... yes
```

Note: If it says "no", check the configuration run output for reasons. Make sure that you have all required build dependencies installed and fixed the EGLFS library references!

#### MesaGL

When building for Raspberry Pi 4, check for

```
  EGLFS .................................. yes
  EGLFS details:
...
    EGLFS EGLDevice ...................... yes
...
    EGLFS Raspberry Pi ................... no
```

### Compile Qt

Now Qt should be configured properly with all features enabled that we need. If you need some of the optional features, make sure to check the configure result that they where properly detected.

To compile Qt run:

```
make
```

or if you are using any of the quad-core Pis, append the *-j4* parameter to build in parallel. Make sure you have proper cooling in this case, the 3+ this should run fine without getting throttled too much.

### Install the build

The compilation should finnish without any errors, if it does not, double check that you have all the dependecies installed and run configure correctly.

If all is well, install Qt by running

```
make install
```

You should now have **Qt 5.12.7** installed in */opt/Qt5.12* ready for use. To configure your Qt project(s) to build with this version run qmake from the installation directory:

```
/opt/Qt5.12/bin/qmake
```

You can of course also add it to your PATH, edit your users "*.profile*" or system wide in *"/etc/profile"* .

## Choosing platform

As explained in the introduction, this build will default to using the eglfs platform, meaning that no windowing environment is required to run GUI applications. You can choose the platform binaries will run against by suppling the "-platform" paramter when running them.
