#test
#環境構築手順として.env.exampleファイルをコピーして .envrc ファイルを生成すること。
#.envrcに環境を構築し、direnv allowコマンドでロードをすること。
#migrationのインストール手順
1."brew install golang-migrate"　このコマンドをCLIに入力
2.migration up の実行
migrate -source file://./migrations/users/ -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up
2.migration down の実行
migrate -source file://./migrations/users/ -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" down
#サーバー起動手順
kensyu/server/main.goにて "go run main.go"を実行
