##
# ttt
#
# @file
# @version 0.1
#
.PHONY: clean

clean:
	rm -vi ttt*

release:
	GOOS=linux GOARCH=amd64 go build -o ttt_linux_amd64
	strip ttt_linux_amd64
	GOOS=windows GOARCH=amd64 go build -o ttt_win_amd64.exe
	strip ttt_win_amd64.exe
	GOOS=darwin GOARCH=amd64 go build -o ttt_osx_amd64
# end
