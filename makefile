all:
	cd public/frontend && npm i && npm run build && cd ../../ && go build --tags=prod cmd/main.go

go:
	go build --tags=prod cmd/main.go