#!/bin/bash
wget https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.19.zip
unzip go_appengine_sdk_linux_amd64-1.9.19.zip
rm go_appengine_sdk_linux_amd64-1.9.19.zip
rm /home/action/.parts/bin/*.py
export PATH=$PATH:/home/action/go_appengine/
cd /home/workspace/go/
git clone https://github.com/alessandro-aglietti/ready-set-go
chmod +x /home/workspace/go/ready-set-go/goapp-serve-nitrous
cd /home/action/.parts/bin
ln -s /home/workspace/go/ready-set-go/goapp-serve-nitrous
cd /home/workspace/go/ready-set-go
