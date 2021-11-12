KUBERNETES_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

##@ Application kubernetes manifests
.PHONY: render-staging render-prod deploy-staging deploy-prod require-tag

render-staging:	## Render helm chart locally using staging values.yaml
	pushd ${KUBERNETES_DIR} && helm template test-release -f ./values.staging.yaml ./chart

deploy-staging: require-tag ## Install helm chart using staging values.yaml
	pushd ${KUBERNETES_DIR} && \
		helm upgrade --install \
					 -f ./values.staging.yaml \
					 --set-string image.tag=${TAG} \
					 --wait \
					 -n test \
					 argocd-playground-app \
					 ./chart

deploy-prod: require-tag ## Install helm chart using production values.yaml
	pushd ${KUBERNETES_DIR} && \
		helm upgrade --install \
					 -f ./values.prod.yaml \
					 --set-string image.tag=${TAG} \
					 --wait \
					 -n test \
					 argocd-playground-app \
					 ./chart

render-prod:	## Render helm chart locally using production values.yaml
	pushd ${KUBERNETES_DIR} && helm template test-release -f ./values.prod.yaml ./chart

require-tag:
ifndef TAG
	$(error TAG variable is required)
endif
