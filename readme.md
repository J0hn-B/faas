# OpenFaas

Create a 2 node K8s cluster with k3d:  
`k3d cluster create dev-cluster --agents 2`

Install ArgoCD:
`cd scripts && ./argocd_install.sh`

Prepare the cluster with Helm Charts:
`kubectl apply -f k8s/argo_config/`

Log into your OpenFaaS gateway:  
`kubectl port-forward svc/gateway -n openfaas 8080:8080`
`kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo`
