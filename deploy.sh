#!/bin/bash

# 初回のみClone、以降はPullする
if cd Junking; then
  git pull;
else
  git clone $1 Junking;
  cd Junking
fi

# 実験時のGCEでdocekr-composeコマンドが使えなかったため、それ用のdocker imageを使用した
# ※デーモン(-d)にしないとCIがいつまでも終わらないので注意

docker-compose build --build-arg GoogleClientID=$2 --build-arg GoogleClientSecret=$3
docker-compose up -d

#docker run \
#--rm -v /var/run/docker.sock:/var/run/docker.sock \
#-v "$PWD:/$PWD" -w="/$PWD" \
#docker/compose:1.22.0 \
#up -d
