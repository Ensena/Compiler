env GOOS=linux GOARCH=amd64 go build -o  builder main.go
docker build -t ensena/cpp-compiler .
docker push ensena/cpp-compiler