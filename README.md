# Weather-api
Weather-api for consulting weather using point coordinates. Written in Golang with Redis

## How to run
This repo uses Docker for running redis. First, you need Docker installed and run
```
docker run --name redis-container -p 6379:6379 -d redis
```
Then inside this repo run:
```
go mod tidy
```
And then for running:
```
go run main.go
```

