# Get version from git hash
git_hash := $(shell git rev-parse --short HEAD || echo 'development')

# project version
# version = $(shell git describe --tags --abbrev=0 || echo 'development')
version = $(shell git tag | sort -V | tail -1 || echo 'development')
# Get current date
current_time = $(shell date +"%Y-%m-%d:T%H:%M:%S")

name:="mtssctl"

# Add linker flags
linker_flags = '-s -X github.com/piqba/mtss-cli/cmd/cli.buildTime=${current_time} -X github.com/piqba/mtss-cli/cmd/cli.versionHash=${git_hash} -X github.com/piqba/mtss-cli/cmd/cli.version=${version}'

# Build binaries for current OS and Linux
.PHONY:
compile:
	@echo "Building binaries..."

	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o ./build/${name}-${version}-linux-amd64 cmd/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags=${linker_flags} -o ./build/${name}-${version}-windows-amd64.exe cmd/main.go

compress:
	./upx -9 -q ./build/${name}-${version}-linux-amd64