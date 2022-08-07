#!/bin/zsh

GOOS=linux GOARCH=amd64 go build -o collector main.go
