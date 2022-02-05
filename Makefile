build-frontend:
	cd frontend && npm run build

build-mac: build-frontend
	GOOS=mac go build -o bin/macos/frontend.exe

build-windows: build-frontend
	GOOS=windows go build -o bin/windows/frontend.exe

server-run-dev:
	GIN_MODE=debug go run server/server.go

server-run-test:
	GIN_MODE=test go run server/server.go

server-run-prod:
	GIN_MODE=release go run server/server.go