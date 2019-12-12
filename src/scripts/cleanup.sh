#!/usr/bin/env bash
echo ">>>>>>>>>>>>>>>>>LENS PLATFORM<<<<<<<<<<<<<<<<<<<<<"
echo "Stopping All Docker Containers"
docker stop $(docker ps -a -q)
echo ">>>>>>>>>>>>>>>>>LENS PLATFORM<<<<<<<<<<<<<<<<<<<<<"
echo "Removing All Stopped Docker Containers"
docker rm $(docker ps -a -q)
echo ">>>>>>>>>>>>>>>>>LENS PLATFORM<<<<<<<<<<<<<<<<<<<<<"
