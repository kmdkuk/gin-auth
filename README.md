# gin-auth

```
docker-compose up -d
# user create
curl -X 'POST' 'http://localhost:3000/users' -H 'accept: */*' -H 'Content-Type: application/json' -d '{"id": "string","password": "string"}'
# user login
curl -v -X 'POST' 'http://localhost:3000/login' -H 'accept: */*' -H 'Content-Type: application/json' -d '{"id": "string", "password": "string" }'
# Set-Cookieの内容をコピー
MYSESSION="コピペした内容"
# MYSESSION=MTYzNDQ1NzQzOHxEdi1CQkFFQ180SUFBUkFCRUFBQUp2LUNBQUVHYzNSeWFXNW5EQWdBQm5WelpYSnBaQVp6ZEhKcGJtY01DQUFHYzNSeWFXNW58rm4-L73uxKp-8TQsHfvIWz-7ybCprjQOifKwJcn2jYY=
curl -X 'GET' 'http://localhost:3000/users' -H 'accept: */*' -H 'Cookie: MYSESSION=${MYSESSION}'
```

### 参考にしたBlog
https://zenn.dev/someone7140/articles/02181927acd040
