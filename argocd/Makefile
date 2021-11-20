ARGOCD_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
ARGOCD_SERVER_URL ?= $(shell minikube ip):32032

##@ ArgoCD
.PHONY: login-argocd-server sync-staging

apply-argocd-app-manifest: ## Apply ArgoCD application and project manifest
	cd ${ARGOCD_DIR} && kubectl apply -f ./manifest

login-argocd-server: ## Login to local ArgoCD server
	$(eval PASS = $(shell kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d))
	@argocd login $(shell minikube ip):32032 --username admin --password ${PASS}

sync-staging: ## Sync staging application
	@pushd ${ARGOCD_DIR} && \
		argocd app sync argocd-playground-app-staging --insecure --server ${ARGOCD_SERVER_URL} --auth-token ${ARGOCD_TOKEN} && \
		argocd app wait argocd-playground-app-staging --insecure --server ${ARGOCD_SERVER_URL} --auth-token ${ARGOCD_TOKEN}
