BUILDPATH=$(CURDIR)
API_NAME=server

build: 
	@echo "Creating Binary ..."
	@go build -mod=vendor -ldflags '-s -w' -o $(BUILDPATH)/build/bin/${API_NAME} cmd/main.go
	@echo "Binary generated in build/bin/${API_NAME}"

test: 
	@echo "Executing tests ..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

coverage: 
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -html=coverfile_out -o coverfile_out.html

.PHONY: test build coverage all clean
