#!/bin/sh
srcPath="server"
pkgFile="main.go"
outputPath="build"
app="server"
output="$outputPath/$app"
src="$srcPath/$pkgFile"

rm -f $output
printf "\nBuilding: $app\n"
time go build -o $output $src
printf "\nBuilt: $app size:"
ls -lah $output | awk '{print $5}'
printf "\nDone building: $app\n\n"
