.. dff-vent documentation master file, created by
   sphinx-quickstart on Thu Apr 23 10:13:45 2020.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.



Installation
====================================

.. toctree::
   :maxdepth: 2
   :caption: Contents:


Install on host:

   1. [On Host] Install go_ and `therecipe's wrapper`_

   .. code-block:: bash
      :linenos:
      
      wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
      tar -C /usr/local -xzf go1.14.2.linux-amd64.tar.gz
      sudo apt-get -y install build-essential libglu1-mesa-dev \
         libpulse-dev libglib2.0-dev
      export GO111MODULE=off
      go get -v github.com/therecipe/qt/cmd/... && \
         $(go env GOPATH)/bin/qtsetup test && \
         $(go env GOPATH)/bin/qtsetup -test=false


   3. [On RP4] Install Raspbian (Tested on Buster)

   .. code-block:: bash
      :linenos:

      export QT_PKG_CONFIG=true
      sudo apt-get --no-install-recommends install \
         libqt*5-dev qt*5-dev qml-module-qtquick-* qt*5-doc-html golang
      go get -v -tags=no_env github.com/therecipe/qt/cmd/...
      $(go env GOPATH)/bin/qtsetup
      
.. _go: https://golang.org/dl/
.. _`therecipe's wrapper`: https://github.com/therecipe/qt/