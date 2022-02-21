# Go (Golang) GOOS and GOARCH
# https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63

# How to cross compile Go with CGO programs for a different OS/Arch
# https://gist.github.com/steeve/6905542

CC=gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -v -o ./memorydb.exe -ldflags="-extld=$CC"