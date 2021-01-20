# monorepo

## Quick Start Guide

#### For Windows machines 
1. Install WSL2 https://docs.microsoft.com/en-us/windows/wsl/
2. Install Docker on Windows workstation https://docs.docker.com/docker-for-windows/install/
3. Install Ubuntu 18.04 from Windows Store
4. Turn on WSL integration in Docker Settings
5. Enable Kubernates in Docker Settings
6. Install k8s ingress controller:
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.41.2/deploy/static/provider/cloud/deploy.yaml
```
7. Install Skaffold
https://skaffold.dev/docs/install/

8. To run the development build:
```
cd monorepo
skaffold dev
```
9. Go to http://localhost/
