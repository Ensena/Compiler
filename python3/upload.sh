env GOOS=linux GOARCH=amd64 go build -o  builder main.go
docker build -t ensena/python .
docker push ensena/python