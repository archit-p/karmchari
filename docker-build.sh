#!/bin/bash

if [ ${USER} != "root" ]
then
	echo "${0}: rerun as root" 1>&2
	exit 1
fi

# stop currently running redis instances
docker rm $(docker stop $(docker ps -a -q --filter ancestor="redis")) 1>&2 2>/dev/null

# stop currently running karmchari instances
docker rm $(docker stop $(docker ps -a -q --filter ancestor="karmchari")) 1>&2 2>/dev/null

# launch a redis container
docker run --name redis-karmchari -d redis

# create a new image
docker build -t karmchari .

# run a container with our image
docker run -p 51463:51463 --name karmchari-prod --link redis-karmchari:redis -d karmchari
