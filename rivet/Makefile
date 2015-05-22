CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
TAG=${TAG:-latest}
COMMIT=`git rev-parse --short HEAD`

all: deps build

deps:
	@godep restore

clean:
	@rm -rf Godeps/_workspace rivet

build:
	@godep go build -a -tags 'netgo' -ldflags "-w -X github.com/ehazlett/rivet/version.GITCOMMIT $(COMMIT) -linkmode external -extldflags -static" .

build-cross:
	@gox -os="linux darwin windows" -arch="386 amd64" -ldflags "-w -X github.com/ehazlett/rivet/version.GITCOMMIT $(COMMIT)" -output "rivet_{{.OS}}-{{.Arch}}"

image: build
	@echo Building image $(TAG)
	@docker build -t ehazlett/rivet:$(TAG) .

release: deps build image
	@docker push ehazlett/rivet:$(TAG)

test: clean 
	@godep go test -v ./...

.PHONY: all deps build clean image test release
