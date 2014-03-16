# setup $GOPATH
export GOPATH := $(CURDIR)/.godeps:$(GOPATH)

# source repo
GOREPO := github.com/wilmoore/go-userdir
VENDOR := $(firstword $(subst :, ,$(GOPATH)))/src/$(GOREPO)

clean:
	@$(RM) -rf example
	@$(RM) -rf $(CURDIR)/.godeps

$(VENDOR):
	@mkdir -p $(VENDOR)

example: clean deps
	@go build example.go

test: clean deps
	@go test -v $(GOREPO)

deps: $(VENDOR)
	@cp userdir*.go $(VENDOR)

