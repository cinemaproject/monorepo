# monorepo

## Quick Start Guide

1. Install Kubernetes and set up a cluster
Try following this guide: https://kubernetes.io/blog/2020/05/21/wsl-docker-kubernetes-on-the-windows-desktop/

This deployment requires NGINX ingress controller. On minikube use
```sh
minikube addons enable ingress
```

2. Install Skaffold
https://skaffold.dev/docs/install/

3. To run the development build:
```
cd monorepo
skaffold dev
```
