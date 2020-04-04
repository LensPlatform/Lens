#!/usr/bin/env bash
echo ">>>>>>>>>>>>>>>LENS<<<<<<<<<<<<<<<"
echo "stopping any running containers"
docker-compose stop
docker-compose down -v
docker system prune -f
docker rmi $(docker images -q)
echo "removing stopped container"
docker-compose rm -f
echo "pulling latest containers from docker hub"
docker-compose pull
docker-compose build --no-cache
echo "spinning up services"
docker-compose up --force-recreate