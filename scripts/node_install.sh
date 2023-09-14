#!/bin/bash

wget https://unofficial-builds.nodejs.org/download/release/v20.6.1/node-v20.6.1-linux-armv6l.tar.gz
tar -xzf node-v20.6.1-linux-armv6l.tar.gz
sudo cp -R node-v20.6.1-linux-armv6l/* /usr/local
rm -rf node-v20.6.1-*
PATH=$PATH:/usr/local/bin