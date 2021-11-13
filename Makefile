.DEFAULT_GOAL := help

SHELL = bash

include app/Makefile
include bootstrap/argocd/Makefile

##@ Help
.PHONY: help
help:  ## Display this help.
	./scripts/list-make-targets.sh $(MAKEFILE_LIST)

##@ Github Actions
.PHONY: act
act:  ## Run github actions workflow locally using act
	act -s MANIFEST_REPO_TOKEN -s GITHUB_TOKEN
