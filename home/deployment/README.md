# home

```bash
# first need to create secret for docker auth in ghcr.io
kubectl create secret docker-registry docker-github-registry --docker-server=ghcr.io --docker-username=ak9024 --docker-password=<github_token> --docker-email=<email>
# apply manifest
kubectl apply -f .
```

```bash
# add ingress and annotate to cert-manager
kubectl create ingress home --rule="adiatma.tech/*=home:4173,tls=adiatma.tech"
kubectl annotate ing/home cert-manager.io/cluster-issuer=letsencrypt-production]
```
