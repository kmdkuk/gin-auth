# gin-auth

## getting start

```
cp .env.sample .env
docker-compose up -d
# user create
curl -X 'POST' 'http://localhost:3000/users' -H 'accept: */*' -H 'Content-Type: application/json' -d '{"user_id": "string","password": "string"}'
# user login
curl -v -X 'POST' 'http://localhost:3000/login' -H 'accept: */*' -H 'Content-Type: application/json' -d '{"user_id": "string", "password": "string" }'
# Set-Cookieの内容をコピー
MYSESSION="コピペした内容"
# MYSESSION=MTYzNDQ1NzQzOHxEdi1CQkFFQ180SUFBUkFCRUFBQUp2LUNBQUVHYzNSeWFXNW5EQWdBQm5WelpYSnBaQVp6ZEhKcGJtY01DQUFHYzNSeWFXNW58rm4-L73uxKp-8TQsHfvIWz-7ybCprjQOifKwJcn2jYY=
curl -X 'GET' 'http://localhost:3000/users' -H 'accept: */*' -H 'Cookie: MYSESSION=${MYSESSION}'
```

## architecture

```
❯ tree -d
.
├── constants       : いろんな場所で共通した定数
├── db              : Databaseとのコネクション
├── docs            : swagger用のyaml置き場
├── handler         : ginに渡すハンドラー実装してる場所(controllerといった方が良いかも)
├── middleware      : 認証確認するためのmiddleware
├── model           : DBに保存しておきたいModel(今回の例ならUser)を定義
├── mysql           : mysqlの設定、初期起動時のscript
│   └── init
└── repository      : DatabaseからModelを取り出す部分を実装している場所

10 directories
```

handlerは、repositoryに依存
repositoryは、Databaseに依存

## TODO

- repositoryをMockにしたら、DBに依存することなくhandlerをテスト
- DBのパスワードとかを、環境変数にする

### 参考にしたBlog
https://zenn.dev/someone7140/articles/02181927acd040
