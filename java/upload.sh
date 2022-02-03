env GOOS=linux GOARCH=amd64 go build -o  builder main.go
docker build -t ensena/java-compiler .
docker push ensena/java-compiler