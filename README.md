[![Actions Status: build](https://github.com/OH-DEER-NO-BEER/Junking/actions/workflows/main_junking.yml/badge.svg)](https://github.com/OH-DEER-NO-BEER/Junking/actions?query=workflow%3A"Deploy%20docker-compose%20to%20GCE")

# Junking

<div align="center">
<img src="https://github.com/OH-DEER-NO-BEER/Junking/blob/main/images/logo.png" width=60% height=60%>
</div>

## What is?

今まで相手が出した手，勝率などを見ることができる心理的要素を強化したじゃんけんアプリ

## System Configuration

<div align="center">
<img src="https://github.com/OH-DEER-NO-BEER/Junking/blob/main/images/gh-backend.png" width=60% height=60%>
</div>

<div align="center">
<img src="https://github.com/OH-DEER-NO-BEER/Junking/blob/main/images/gh-front.png" width=60% height=60%>
</div>

<div align="center">
<img src="https://github.com/OH-DEER-NO-BEER/Junking/blob/main/images/gh-infra.png" width=60% height=60%>
</div>

## For Developers
### Deploy
`main`ブランチに`push`されると自動でビルド，本番環境にデプロイされるようになっています．
### docker-compose command
#### build
`$ docker-compose build --build-arg GoogleClientID=ClientID --build-arg GoogleClientSecret=ClientSecret`
#### run
`$ docker-compose up -d`
#### stop
`$ docker-compose down -v`
### container command
#### run bash in the container
`$ docker exec -it <container_name> bash`
#### logout from the container
`$ exit`
  
### Vue.js Building
`$ cd ${PATH_TO_PACKAGE_JSON_FOLDER}`  
`$ npm run build`
