# 短链接项目

项目启动
```go
go run shortener.go
```

生成model文件
```go
goctl model mysql datasource -url="cc:123@tcp(127.0.0.1:3306)/url" -table="short_url_map" -dir="./model" -c
```