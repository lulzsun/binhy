#!/bin/bash

VERSION=1.20

## Download the latest version of Golang
wget https://dl.google.com/go/go$VERSION.linux-armv7l.tar.gz

## Extract the archive
tar -C /usr/local -xzf go$VERSION.linux-armv7l.tar.gz

## Set the environment variable for Golang
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

## Verify the installation
go version

## Clean up
rm go$VERSION.linux-armv7l.tar.gz