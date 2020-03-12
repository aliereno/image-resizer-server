#!/bin/sh
srcPath="server"
pkgFile="main.go"
src="$srcPath/$pkgFile"

printf "\nServer Started running\n"
time go run $src
printf "\nServer Stopped running\n\n"
