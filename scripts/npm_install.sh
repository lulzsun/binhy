#!/bin/bash

# Update local package list
sudo apt-get update

# Install NodeJS and NPM
sudo apt-get install -y nodejs npm

# Update NodeJS and NPM to the latest version
sudo npm install -g n
sudo n stable
sudo npm install -g npm

# Reload PATH
PATH="$PATH"