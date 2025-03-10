OUTPUT_DIR = bin
GO = go
GOFLAGS = -v -ldflags="-s -w"

gen-api:
	goctl api go --api api/opensql.api --dir . --style go_zero

build-linux-amd64:
	@echo "compile Linux amd64"
	@set GOOS=linux&& set GOARCH=amd64&& ${GO} build $(GOFLAGS) -o $(OUTPUT_DIR)/linux-amd64/$(BINARY_NAME) .