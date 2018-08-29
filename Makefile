all: windows-amd64 darwin-amd64 linux-amd64

windows-amd64: gocourse-windows-amd64/gocourse-windows-amd64.exe gocourse-windows-amd64/profile.json gocourse-windows-amd64/run.bat

darwin-amd64: gocourse-darwin-amd64/gocourse-darwin-amd64 gocourse-darwin-amd64/profile.json

linux-amd64: gocourse-linux-amd64/gocourse-linux-amd64 gocourse-linux-amd64/profile.json


%/profile.json: profile.json.template
	cp -f $< $@

gocourse-windows-amd64/run.bat: run.bat
	cp -f $< $@

gocourse-windows-amd64/gocourse-windows-amd64.exe: main.go
	GOOS=windows GOARCH=amd64 go build -o $@ $<

gocourse-darwin-amd64/gocourse-darwin-amd64: main.go
	GOOS=darwin GOARCH=amd64 go build -o $@ $<

gocourse-linux-amd64/gocourse-linux-amd64: main.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

.PHONY: clean

clean:
	rm -rf gocourse-windows-amd64
	rm -rf gocourse-darwin-amd64
	rm -rf gocourse-linux-amd64