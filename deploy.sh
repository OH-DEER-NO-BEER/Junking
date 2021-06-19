#!/bin/bash

# 初回のみClone、以降はPullする
if cd Junking; then
  git pull;
else
  git clone $1 Junking;
  cd Junking
fi

docker-compose build --no-cache --build-arg GoogleClientID=$2 --build-arg GoogleClientSecret=$3
docker-compose up -d
