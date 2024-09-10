#!/bin/sh
go build -o helper
chmod +x ./helper
sudo cp ./helper /usr/local/bin/helper
echo "New helper version installed"