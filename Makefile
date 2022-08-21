VERSION=1.0.0
APPNAME=GoMarkdown

get-fyne:
	@go get -u fyne.io/fyne/v2

install-fyne-cmd:
	@go install fyne.io/fyne/v2/cmd/fyne@latest

build: clean
	@go build -o gomarkdown .

clean:
	@go clean
	@rm -f gomarkdown

package:
	@rm -rf GoMarkdown.app
	@fyne package -appVersion ${VERSION} -name ${APPNAME} -release

test:
	@go test -v ./...

.PHONY: build clean package test get-fyne install-fyne-cmd