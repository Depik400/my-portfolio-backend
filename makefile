all:
	cd public/frontend && npm i && npm run build && cd ../../ && go build --tags=prod cmd/main.go

go:
	go build --tags=prod cmd/main.go

macos:
	set GOOS=darwin&& set GOARCH=arm64&& go build --tags=prod -o macos .\cmd\main.go