# Junking


## for developers
### Docker-composeのコマンド
- コンテナビルド : docker-compose build --build-arg GoogleClientID=別途共有のClientID --build-arg GoogleClientSecret=別途共有のClientSecret
- コンテナ起動(サーバーまで自動で起動してくれるはず) : docker-compose up -d
- コンテナ＆ボリューム削除 : docker-compose down -v
### コンテナ内コマンド
- コンテナログイン : docker exec -it bash
- サーバ立ち上げ : go run server.go
- サーバー終了 : Ctrl-c
- コンテナログアウト : exit