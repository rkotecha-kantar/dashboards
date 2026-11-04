.PHONY: build-frontend build run dev-frontend test

build-frontend:
	cd frontend && npm run build

run: build-frontend
	cd backend && go run ./cmd/server

build: build-frontend
	cd backend && go build -o ../bin/server ./cmd/server

dev-frontend:
	cd frontend && npm run dev

test:
	cd backend && go test ./...