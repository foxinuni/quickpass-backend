generate-wire:
	@wire ./...

generate: generate-wire

build: generate
	@go build -o build/ ./...

start: build
	@./build/backend.exe