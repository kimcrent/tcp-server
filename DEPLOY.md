# Docker + GitHub Actions deploy

## 1) SSH deploy key

```bash
# Generated key path:
~/.ssh/github_actions_go_deploy

# If you need to regenerate:
ssh-keygen -t ed25519 -C "github-actions-go-deploy" -f ~/.ssh/github_actions_go_deploy -N ""
```

## 2) Add GitHub repository secrets

- `SERVER_HOST` = `82.40.38.98`
- `SERVER_USER` = `k1epa`
- `SERVER_SSH_KEY` = contents of `~/.ssh/github_actions_go_deploy`
- `GHCR_USERNAME` = your GitHub username
- `GHCR_TOKEN` = GitHub PAT with `read:packages` (for server pull from GHCR)

Quick command to print private key for `SERVER_SSH_KEY`:

```bash
cat ~/.ssh/github_actions_go_deploy
```

## 3) How deploy works

- Push to `main` triggers `.github/workflows/deploy.yml`
- CI builds and pushes image to GHCR:
  - `ghcr.io/<owner>/go-tcp-server:latest`
  - `ghcr.io/<owner>/go-tcp-server:<commit_sha>`
- CD logs into the server, updates `/opt/go-tcp-server/docker-compose.yml`, pulls new image and restarts container.

## 4) Manual check on server

```bash
ssh k1epa@82.40.38.98
cd /opt/go-tcp-server
docker compose ps
docker logs --tail=100 go-tcp-server
```
