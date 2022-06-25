include .env
export


.PHONY: proto
proto:
	protoc	--go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative ./infrastructure/grpc/proto/cryptocore.proto

.PHONY: build_linux
build_linux:
	env GOOS=linux GOARCH=amd64 go build -o cryptocore

.PHONY: build_mac
build_mac:
	env GOOS=darwin GOARCH=arm64 go build -o cryptocore

.PHONY: build_windows
build_windows:
	env GOOS=windows GOARCH=amd64 go build -o cryptocore