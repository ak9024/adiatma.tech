# api

```bash
# first need to create secret for docker auth in ghcr.io
kubectl create secret docker-registry docker-github-registry --docker-server=ghcr.io --docker-username=ak9024 --docker-password=<github_token> --docker-email=<email>
# apply manifest
kubectl apply -f .
```

```bash
# add ingress and annotate to cert-manager
kubectl create ingress api --rule="api.adiatma.tech/*=api:8000,tls=api.adiatma.tech"
kubectl annotate ing/api cert-manager.io/cluster-issuer=letsencrypt-production
```
