.PHONY: minikube-manifest install uninstall server-url admin-password

minikube-manifest:
	kustomize build ./install/overlays/minikube

install:
	kustomize build ./install/overlays/minikube | kubectl apply -f -

uninstall:
	kustomize build ./install/overlays/minikube | kubectl delete -f -

server-url:
	echo "https://$(shell minikube ip):32032"

admin-password:
	kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
