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
`faas-cli login --username admin --password your_password_here`

> Functions can be invoked via a GET or POST method only.

faas-cli new --lang dockerfile long-task --prefix="johnb79"

faas-cli up -f hello-openfaas.yml # build --> push --> deploy, in one command

faas-cli new --lang python3 astronaut-finder --prefix="johnb79"

faas-cli build -f ./astronaut-finder.yml

faas-cli push -f ./astronaut-finder.yml

faas-cli up -f astronaut-finder.yml

kubectl scale deployment --replicas=0 long-task -n openfaas-fn

faas-cli login --username admin --password HTDn2zETtfjt
