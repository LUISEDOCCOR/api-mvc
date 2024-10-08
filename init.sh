#!/bin/bash

echo "Start MCV-API"

go mod tidy

case $1 in
    --dev)
        go install github.com/air-verse/air@latest
        export PATH=$PATH:$HOME/go/bin
        clear
        air
    ;;
    --prod)
        clear
        go build -o ./build/main
        ./build/main
    ;;
    --docs)
        rm -rf ./docs/*
        tree > ./docs/tree.txt
        sqlite3 gorm.sqlite .dump > ./docs/gorm.sql
    ;;
    *)
        echo "Invalid --tag"
    ;;
esac
