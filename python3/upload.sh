env GOOS=linux GOARCH=amd64 go build -o  builder cmd/main.go
docker build -t elmalba/ensena-python-test .
docker push elmalba/ensena-python-test