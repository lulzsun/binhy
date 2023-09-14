#!/bin/bash

wget https://unofficial-builds.nodejs.org/download/release/v18.12.0/node-v18.12.0-linux-armv6l.tar.gz
tar xvfJ node-v18.12.0-linux-armv6l.tar.gz
sudo cp -R node-v18.12.0-linux-armv6l/* /usr/local
rm -rf node-v18.12.0-*
PATH=$PATH:/usr/local/bin