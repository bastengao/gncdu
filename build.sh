GOOS=linux GOARCH=amd64 go build -o gncdu-linux-amd64
GOOS=darwin GOARCH=amd64 go build -o gncdu-darwin-amd64
GOOS=windows GOARCH=amd64 go build -o gncdu-windows-amd64.exe
GOOS=windows GOARCH=386 go build -o gncdu-windows-386.exe