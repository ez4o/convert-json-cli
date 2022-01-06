release-windows:
	@go build -o .\convjson.exe

release-linux:
	@REM avoid white space after env variable
	@set CGO_ENABLED=0&& set GOOS=linux&& set GOARCH=amd64&& go build -o .\convjson

release-all:
	@make release-windows
	@make release-linux