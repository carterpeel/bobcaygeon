#!/bin/bash

wget https://archive.raspbian.org/raspbian.public.key -O - | sudo apt-key add -
sudo apt-get update
sudo apt-get install libasound2-dev

# all the gneration was handled when doing the x86 linux build, so we 
# just build the go code

echo "Starting bcg build"
go build -mod=vendor -o bcg-arm cmd/bcg.go
go build -mod=vendor -o bcg-mgmt-arm cmd/mgmt/bcg-mgmt.go
