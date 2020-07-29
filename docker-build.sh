#!/bin/bash

if [ ${USER} != "root" ]
then
	echo "${0}: rerun as root" 1>&2
	exit 1
fi

# create a new image
docker build -t karmchari .

# run a container with our image
docker run -p 51463:51463 -it karmchari
