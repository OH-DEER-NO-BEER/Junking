#!/bin/bash

# 初回のみClone、以降はPullする
if cd Junking; then
  git pull;
else
  git clone $1 Junking;
  cd Junking
fi

docker-compose down -v
docker-compose build --build-arg GoogleClientID=$2 --build-arg GoogleClientSecret=$3 --build-arg CertFile=$4 --build-arg KeyFile=$5
docker-compose up -d
