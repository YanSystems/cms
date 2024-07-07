#!/bin/bash

if [ "$(docker ps -a -q -f name=yan-cms)" ]; then
    echo "Using existing container yan-cms."
    docker start yan-cms
else
    echo "Creating and starting new container yan-cms."
    docker run --name yan-cms -d cms
fi