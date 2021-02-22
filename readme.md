# OpenFaas playground

Create an Openfaas gitops playground for testing locally with:  
[k3d](https://k3d.io/)
ArgoCD
Openfaas

Create a 3 node K8s cluster with k3d:  
`k3d cluster create multiserver --servers 3`

Install ArgoCD:
`cd scripts && ./argocd_install.sh`

Prepare the cluster with Helm Charts:
`kubectl apply -f k8s/argo_config/`

Log into your OpenFaaS gateway:  
`kubectl port-forward svc/gateway -n openfaas 8080:8080`
`kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo`
`faas-cli login --username admin --password your_password_here`

## Faas How to

> Functions can be invoked via a GET or POST method only.

faas-cli new --lang dockerfile long-task --prefix="your_docker_registry_here"

faas-cli up -f hello-openfaas.yml # build --> push --> deploy, in one command

## Example

faas-cli new --lang python3 astronaut-finder --prefix="your_docker_registry_here"

faas-cli build -f ./astronaut-finder.yml

faas-cli push -f ./astronaut-finder.yml

faas-cli up -f astronaut-finder.yml
