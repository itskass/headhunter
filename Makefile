# installs HEADHunter on the local machine as
# hhunt
.phony: install
install:
	go build -i -o hhunt .

# builds cross platform distributions for windows and 
# linux. mac build not working, run `make _mac` on a 
# macOS system instead.
.phony: builds
builds: 
		docker run --rm -v \
		"$$PWD":/go/src/github.com/itskass/headhunter \
		-w /go/src/github.com/itskass/headhunter \
		golang:1.13 make _linux _windows

.phony: _linux
_linux: 
	go get -d -v ./...
	go get -d github.com/btcsuite/btcd/btcec 
	go build -o dist/hhunt_linux .

.phony: _windows
_windows:
	go get -d github.com/pkg/errors gopkg.in/natefinch/npipe.v2
	env GOOS=windows GOARCH=386 go build -o dist/hhunt_windows.exe .

.phony: _mac
_mac:
	nv GOOS=darwin GOARCH=amd64 go build -o dist/hhunt_mac .
