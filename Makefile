##
# ttt - time tracker tool
#
# @file
# @version 0.1
.PHONY: release clean test benchmark

nixbuild:
	nix-build -E 'with import <nixpkgs> {}; callPackage ./default.nix {}'

debug:
	CGO_ENABLED=0 go build -o ./bin/ttt_debug

release: linux windows osx

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ttt_linux_amd64

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/ttt_win_amd64.exe

osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/ttt_osx_amd64

test:
	go vet .
	staticcheck .
	govulncheck .
	revive

#benchmark:
#	go test -bench=. -benchmem

clean:
	rm -v ./bin/ttt_* *.pdf

# end
