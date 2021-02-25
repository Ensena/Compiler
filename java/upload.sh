env GOOS=linux GOARCH=amd64 go build -o  builder cmd/main.go
docker build -t elmalba/ensena-java-compiler .
docker push elmalba/ensena-java-compiler