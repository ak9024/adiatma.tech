<div align="center">
    <img src="./assets/architecture.png" />
    <h2>https://adiatma.tech</h2>
    <small>My personal site</small>
</div>

### Intro

Hello, for learning purpose i build this site.

> Deprecate Caddyfile and replace with cert-manager (TLS) and traefik (Load Balancer)

### Prerequisite

- Docker

### Getting Started

```bash
# pull docker images from ghcr.io
docker compose pull

# export env
export PORT=8000
export BASE_PATH=/
export HOST=localhost:8000

# run the container
docker compose up -d
```

### Stack

- Go to serve API's
- Vanilla JS as a static site for home

### API's

Visit in `https://api.adiatma.tech/`

| Endpoints                            | Description |
| ------------------------------------ | ----------- |
| /metrics                         |             |
| /hadith/:author?page=1&perPage=5 |             |
| /list/authors                    |             |
| /swagger/index.html              |             |

### Home

Visit in `https://adiatma.tech`

### License

- [MIT](./LICENSE)

### Roadmap

- [x] API Hadith
- [ ] API Blogs (get data from notion, dev.to)
- [ ] Revamp Home
- [ ] CLI to generate auto change's log
- [ ] Dev `https://playground.adiatma.tech`
