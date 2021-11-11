.DEFAULT_GOAL := help

SHELL = bash

include app/Makefile
include kubernetes/app/Makefile
include kubernetes/argocd/Makefile

##@ Help
.PHONY: help
help:  ## Display this help.
	./scripts/list-make-targets.sh $(MAKEFILE_LIST)
