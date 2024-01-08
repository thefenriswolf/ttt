##
# ttt - time tracker tool
#
# @file
# @version 0.1
.PHONY: release clean test benchmark

build:
	CGO_ENABLED=0 go build -o ttt

release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ttt_linux_amd64
	strip ttt_linux_amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ttt_osx_amd64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ttt_win_amd64.exe
	strip ttt_win_amd64.exe

test:
	go test -v

benchmark:
	go test -bench=. -benchmem

clean:
	rm -v ttt_linux_* ttt_osx_* ttt_win_* ttt

# end
