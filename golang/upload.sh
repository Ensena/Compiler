env GOOS=linux GOARCH=amd64 go build -o  builder main.go
docker build -t ensena/go-compiler .
docker push ensena/go-compiler