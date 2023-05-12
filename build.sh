set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=windows
go build -ldflags "-w -s" -o fscan main.go




CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -ldflags "-w -s" -o fscan.exe main.go