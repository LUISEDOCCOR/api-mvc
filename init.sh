#!/bin/bash

echo "Start MCV-API"

case $1 in
    --dev)
        go install github.com/air-verse/air@latest
        export PATH=$PATH:$HOME/go/bin
        air
        ;;
    --prod)
        go run main.go
        ;;

esac
