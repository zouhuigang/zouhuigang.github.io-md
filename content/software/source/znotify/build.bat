@echo off
set GOOS=linux
set GOPACH=amd64
go build -o ZnotifyServer
pause
