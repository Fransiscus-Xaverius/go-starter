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