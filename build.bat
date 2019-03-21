set GOOS=windows&& set GOARCH=386&& go build -o plamienokWin.exe server/cmd/main.go
set GOOS=linux&& set GOARCH=386&& go build -o plamienokLinux server/cmd/main.go
set GOOS=darwin&& set GOARCH=amd64&& go build -o plamienokMac server/cmd/main.go