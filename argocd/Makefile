ARGOCD_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

##@ ArgoCD
.PHONY: login-argocd-server sync-staging

login-argocd-server: ## Login to local ArgoCD server
	$(eval PASS = $(shell kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d))
	@argocd login $(shell minikube ip):32032 --username admin --password ${PASS}

sync-staging: ## Sync staging application
	pushd ${ARGOCD_DIR} && \
		argocd app sync argocd-playground-app-staging && \
		argocd app wait argocd-playground-app-staging
