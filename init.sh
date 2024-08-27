#!/bin/bash

echo "Start MCV-API"

case $1 in
    --dev)
        go install github.com/air-verse/air@latest
        export PATH=$PATH:$HOME/go/bin
        clear
        air
        ;;
    --prod)
        clear
        go build -o main
        ./main
        ;;
    *)
        echo "Invalid --tag"
    ;;
esac
