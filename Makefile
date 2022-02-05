frontend-build:
	cd frontend && npm run build

server-build-mac: build-frontend
	GOOS=mac go build -o bin/macos/frontend.exe

server-build-windows: build-frontend
	GOOS=windows go build -o bin/windows/frontend.exe

server-run-dev:
	GIN_MODE=debug gin run server/server.go

server-run-test:
	GIN_MODE=test go run server/server.go

server-run-prod:
	GIN_MODE=release go run server/server.go