# Junking

<div align="center">
<img src="https://github.com/OH-DEER-NO-BEER/Junking/blob/main/images/logo.png" width=60% height=60%>
</div>

## What is?

今まで相手が出した手，勝率などを見ることができる心理的要素を強化したじゃんけんアプリ

## システム構成

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
### デプロイ
`main`ブランチに`push`されると自動でビルド，本番環境にデプロイされるようになっています．
### Docker-composeのコマンド
#### コンテナビルド 
`$ docker-compose build --build-arg GoogleClientID=ClientID --build-arg GoogleClientSecret=ClientSecret`
#### コンテナ起動 
`$ docker-compose up -d`
#### コンテナ＆ボリューム削除
`$ docker-compose down -v`
### コンテナ内コマンド
#### コンテナログイン 
`$ docker exec -it [image_name] bash`
#### コンテナログアウト
`$ exit`
