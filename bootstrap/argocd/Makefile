BOOTSTRAP_ARGOCD_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

##@ Bootstrap ArgoCD
.PHONY: minikube-manifest install uninstall server-url admin-password

minikube-manifest: ## Kustomize build to generate manifest for ArgoCD installation
	pushd ${BOOTSTRAP_ARGOCD_DIR} && kustomize build ./overlays/minikube

install: ## Install ArgoCD using kustomize
	pushd ${BOOTSTRAP_ARGOCD_DIR} && kustomize build ./overlays/minikube | kubectl apply -f -

uninstall: ## Uninstall ArgoCD
	pushd ${BOOTSTRAP_ARGOCD_DIR} && kustomize build ./overlays/minikube | kubectl delete -f -

server-url: ## Show URL to ArgoCD server
	@echo "https://$(shell minikube ip):32032"

admin-password: ## Show password for admin user
	@kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

generate-token: ## Generate token for account `github-actions`
	argocd account generate-token -a github-actions
