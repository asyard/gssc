# gssc

[![Go Report Card](https://goreportcard.com/badge/github.com/asyard/gssc)](https://goreportcard.com/report/github.com/asyard/gssc)


![gssc](https://github.com/asyard/gssc/blob/master/img/gssc.png?raw=true "gssc")



**Simple, go-based, unofficial CLI Tool For [Galatasaray Sozluk](https://rerererarara.net/)**


## Usage
Type gssc-xxx (bugun | gundem | yazar | entry)
> gssc-x bugun

![](https://github.com/asyard/gssc/blob/master/img/gssg.gif)



### Build For Your Own
After obtaining source code (and go environment is installed), you may type one of these from your terminal :

for windows (64-bit)

`GOOS=windows GOARCH=amd64 go build -o bin/gssc-v0.1.0-beta.1-windows-amd64.exe`

for mac (64-bit)

`GOOS=darwin GOARCH=amd64 go build -o bin/gssc-v0.1.0-beta.1-darwin-amd64`

`GOOS=darwin GOARCH=arm64 go build -o bin/gssc-v0.1.0-beta.1-darwin-arm64`

for linux (64-bit)

`GOOS=linux GOARCH=amd64 go build -o bin/gssc-v0.1.0-beta.1-linux-amd64`

for custom, first type `go tool dist list`. The result will be similar to (GOOS/GOARCH):

>`...
linux/386
linux/amd64
linux/arm
linux/arm64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
windows/386
windows/amd64
windows/arm
windows/arm64

your local GOOS and GOARCH can be obtained like : 

`go env GOOS GOARCH`

Directly used libraries are cobra v1.5.0[^1], promptui v0.9.0[^2]


[^1]: Cobra https://github.com/spf13/cobra/  
[^2]: promptui https://github.com/manifoldco/promptui