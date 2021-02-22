# OpenFaas playground

Create an Openfaas gitops playground for testing locally with:  
[k3d](https://k3d.io/)  
[ArgoCD](https://argo-cd.readthedocs.io/en/stable/)  
[Openfaas](https://www.openfaas.com/)

Create a 3 node K8s cluster with k3d:  
`k3d cluster create multiserver --servers 1 --agents 2`

Install ArgoCD:  
`cd scripts && ./argocd_install.sh` --> Follow the script and login to ArgoCD UI.

Prepare the cluster with Helm Charts:  
`kubectl apply -f k8s/argo_config/` --> App of Apps pattern, check ArgoCD UI

Wait for OpenFaas to complete installation.  
`kubectl get deployment gateway -n openfaas --watch`

> If gateway deployment fails, delete the gateway deployment and re-sync Argo

Log into your OpenFaaS gateway:  
`kubectl port-forward svc/gateway -n openfaas 8080:8080`

`PASS=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode;)`

`faas-cli login --username admin --password $PASS`

## OpenFaas How to

> Functions can be invoked via a GET or POST method only.

// create new function:  
faas-cli new --lang dockerfile long-task --prefix="your_docker_registry_here"

// build and deploy function:  
faas-cli up -f hello-openfaas.yml # build --> push --> deploy, in one command

## Deploy an example function

A basic function, using the go template.  
It returns time, date and the hostname of the container.

`cd faas/functions`

`faas-cli up -f first.yml`

Check your OpenFaas Web UI. Invoke function and check the result.

`kubectl get deployment first -n openfaas-fn`
