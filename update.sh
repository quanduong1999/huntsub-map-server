#!/bin/sh

pm2 stop qrcode-qrcode-public-api-server
rm qrcode-qrcode-public-api-server.exe
go build -o qrcode-qrcode-public-api-server.exe
pm2 start qrcode-qrcode-public-api-server.exe

while true; do
    
    if [[ $(git pull origin master) == *up-to-date* ]]; 
    then
        echo "no change"
    else
        echo "detect changes"
        sleep 2s
        echo "stop app"
		pm2 stop qrcode-qrcode-public-api-server
		rm qrcode-qrcode-public-api-server.exe
        go build -o qrcode-qrcode-public-api-server.exe
        pm2 restart qrcode-qrcode-public-api-server
    fi

    echo "sleep 30s"

    sleep 30s        

done
