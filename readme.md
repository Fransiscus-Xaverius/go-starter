# How to start

```shell
git clone  https://github.com/Fransiscus-Xaverius/go-starter.git . go-example

docker compose up -d

go mod tidy

go run main.go
```

## Mock generator

```shell
go install github.com/golang/mock/mockgen@latest

# vim ~/.zshrc , add line 
export PATH=$PATH:$(go env GOPATH)/bin

# after update ~/.zshrc, then run
source ~/.zshrc

# check mockgen executable
which mockgen
```

## Docker

```shell
docker compose up -d

docker build -t demo:v1.0.0 .

docker run --name demo -e MYSQL_HOST=host.docker.internal -p 4000:3000 -d demo:v1.0.0
```