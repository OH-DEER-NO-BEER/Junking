#!/bin/bash

# 初回のみClone、以降はPullする
if cd app; then
  git pull;
else
  git clone $1 app;
  cd app
fi

# 実験時のGCEでdocekr-composeコマンドが使えなかったため、それ用のdocker imageを使用した
# ※デーモン(-d)にしないとCIがいつまでも終わらないので注意
docker run \
--rm -v /var/run/docker.sock:/var/run/docker.sock \
-v "$PWD:/$PWD" -w="/$PWD" \
docker/compose:1.22.0 \
up -d
