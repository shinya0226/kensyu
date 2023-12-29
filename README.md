# kensyu
#test
#環境構築手順として.env.exampleファイルをコピーして .envrc ファイルを生成すること。
#fixtureにてテストを実行する際はfix_testデータベースにてusersテーブルを使用する。
#.envrcに環境を構築し、direnv allowコマンドでロードをすること。
#migrationのインストール手順
1.brew install golang-migrate　このコマンドをCLIに入力
2.migration up の実行
migrate -source file://./migrations/users/ -database 'mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HPST}:${DB_PORT})/${DB_NAME}' up
2.migration down の実行
migrate -source file://./migrations/users/ -database 'mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HPST}:${DB_PORT})/${DB_NAME}' down