#!/bin/bash

wget https://nodejs.org/dist/v20.6.1/node-v20.6.1-linux-armv7l.tar.xz
tar -xf node-v20.6.1-linux-armv7l.tar.gz
sudo cp -R node-v20.6.1-linux-armv7l/* /usr/local
rm -rf node-v20.6.1-*
PATH=$PATH:/usr/local/bin