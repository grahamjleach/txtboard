VERSION     := $(shell git rev-parse --short=8 HEAD)
HOSTGOPATH  := $(GOPATH)
LOCALGOPATH := $(CURDIR)/vendor:$(GOPATH)

export GOPATH=$(LOCALGOPATH)

.PHONY: help
help:
	$(info Available targets: )
	$(info | help          )
	$(info | build         )
	$(info | clean         )
	$(info | vet           )
	$(info | lint          )
	$(info | fmt           )
	$(info | quicktest     )
	$(info | test          )
	$(info | printvars     )
	@exit 0

##########################################################################################
## Main targets
##########################################################################################

.DEFAULT_GOAL := build

.PHONY: build
build:
	@CGO_ENABLED=0 GOOS=darwin go build -o bin/txtboard -ldflags "-s -X main.version=$(VERSION)" -a .;

.PHONY: dep
dep:
	@dep ensure -update;

##########################################################################################
## Make utilities
##########################################################################################

.PHONY: printvars
printvars:
	@$(foreach V, $(sort $(.VARIABLES)), $(if $(filter-out environment% default automatic, $(origin $V)), $(warning $V=$($V) )))
	@exit 0
