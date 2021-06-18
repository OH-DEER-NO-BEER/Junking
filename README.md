# Junking


## for developers
### コンテナ
- コンテナ起動 : docker-compose up -d --build
- コンテナ＆ボリューム削除 : docker-compose down -v
### サーバー起動
- コンテナログイン : docker exec -it bash
- サーバー立ち上げ : go run server.go (ただし /app/lib/google/main.go のClientID, ClientSecret は別途共有）
- サーバー終了 : Ctrl-c
- コンテナログアウト : exit