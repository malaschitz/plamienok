#!/usr/bin/env bash

GOOS=windows GOARCH=386 go build -o plamienokWin.exe server/cmd/main.go
GOOS=linux GOARCH=386 go build -o plamienokLinux server/cmd/main.go
GOOS=darwin GOARCH=amd64 go build -o plamienokMac server/cmd/main.go