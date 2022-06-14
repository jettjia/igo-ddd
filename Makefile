GO111MODULES=on
APP=gin-ddd

.PHONY: all
tidy:
	$(eval files=$(shell find . -name go.mod))
	@set -e; \
	for file in ${files}; do \
		goModPath=$$(dirname $$file); \
		cd $$goModPath; \
		go mod tidy; \
		cd -; \
	done

run:
	go run main.go

build:
	set CGO_ENABLED=0
	set GOOS=linux
	set GOARCH=amd64
	go build -o ./bin/linux_amd64/${APP} main.go

test:
	go test -cover ./...