##
# ttt - time tracker tool
#
# @file
# @version 0.1
.PHONY: release clean test benchmark

debug:
	CGO_ENABLED=0 go build -o ttt

release: linux windows osx

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ttt_linux_amd64

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ttt_win_amd64.exe

osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ttt_osx_amd64

test:
	go vet .
	staticcheck .

#benchmark:
#	go test -bench=. -benchmem

clean:
	rm -v ttt_linux_* ttt_osx_* ttt_win_* ttt

# end
