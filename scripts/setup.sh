#!/bin/bash

set -e

apt-get update && apt-get install -y vlc

chmod +x ./go_install.sh
chmod +x ./node_install.sh

./go_install.sh
./node_install.sh

source ~/.bashrc