#!/usr/bin/env bash

DOCKDIR=`dirname $0`
TAG=gohttp

docker build -t demoliang.$TAG --build-arg VERSION=$VERSION --build-arg HTTPPORT=$HTTPPORT --no-cache=true $DOCKDIR
if [ $? -ne 0 ] ;then
	echo "build failed"
	exit 1
fi
