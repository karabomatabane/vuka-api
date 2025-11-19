# Deployment Guide

This guide explains how to deploy the Vuka API to production using GitHub Actions.

## Prerequisites

### 1. GitHub Secrets

You need to add the following secrets to your GitHub repository:

1. Go to your repository on GitHub
2. Navigate to `Settings` → `Secrets and variables` → `Actions`
3. Add the following secret:

- **`SSH_PRIVATE_KEY`**: Your SSH private key that has access to the production server

To generate or use an existing SSH key:

```bash
# If you need to generate a new key pair
ssh-keygen -t ed25519 -C "github-actions@vuka-api" -f ~/.ssh/vuka_deploy

# Copy the private key content (this goes into GitHub Secrets)
cat ~/.ssh/vuka_deploy

# Copy the public key to your server
ssh-copy-id -p 2222 -i ~/.ssh/vuka_deploy.pub yetu@41.71.105.58
```

### 2. Server Setup

Ensure your production server has:

1. **Directory structure**:
   ```bash
   sudo mkdir -p /opt/vuka
   sudo chown yetu:yetu /opt/vuka
   ```

2. **Systemd service file** at `/etc/systemd/system/vuka-api.service`:
   ```ini
   [Unit]
   Description=Vuka API Service
   After=network.target postgresql.service

   [Service]
   Type=simple
   User=yetu
   Group=yetu
   WorkingDirectory=/opt/vuka
   ExecStart=/opt/vuka/vuka-api
   Restart=always
   RestartSec=5
   StandardOutput=journal
   StandardError=journal

   # Environment variables
   Environment="PORT=8080"
   EnvironmentFile=-/opt/vuka/.env

   [Install]
   WantedBy=multi-user.target
   ```

3. **Enable the service**:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable vuka-api
   ```

4. **Environment file** at `/opt/vuka/.env`:
   ```bash
   sudo nano /opt/vuka/.env
   # Add your production environment variables
   ```

## Deployment Process

### Automatic Deployment (Recommended)

The deployment happens automatically when you push a version tag to the main branch:

1. **Merge your changes to main**:
   ```bash
   git checkout main
   git merge develop
   git push origin main
   ```

2. **Create and push a version tag**:
   ```bash
   # Using the helper script (recommended)
   ./scripts/release.sh 1.2.3
   
   # Or manually
   git tag -a v1.2.3 -m "Release v1.2.3"
   git push origin v1.2.3
   ```

3. **Monitor the deployment**:
   - Go to the `Actions` tab in your GitHub repository
   - Watch the "Deploy to Production" workflow
   - The deployment includes automatic verification

### What Happens During Deployment

1. Checks out the tagged code
2. Sets up Go environment
3. Builds the binary with optimizations (`CGO_ENABLED=0 GOOS=linux GOARCH=amd64`)
4. Sets correct permissions (750)
5. Connects to server via SSH (port 2222)
6. Stops the running service
7. Backs up the current binary
8. Deploys the new binary
9. Starts the service
10. Verifies the service is running

## Version Tagging Convention

Follow semantic versioning: `vMAJOR.MINOR.PATCH`

- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes

Examples:
- `v1.0.0` - Initial release
- `v1.1.0` - Added pagination feature
- `v1.1.1` - Fixed pagination bug
- `v2.0.0` - Breaking API changes

## Troubleshooting

### Deployment Failed

1. **Check GitHub Actions logs**:
   - Go to Actions tab → Select the failed workflow → View logs

2. **Check server logs**:
   ```bash
   ssh -p 2222 yetu@41.71.105.58
   sudo journalctl -u vuka-api -n 100 -f
   ```

3. **Check service status**:
   ```bash
   ssh -p 2222 yetu@41.71.105.58
   sudo systemctl status vuka-api
   ```

### Rollback to Previous Version

If deployment fails, the system keeps backups:

```bash
ssh -p 2222 yetu@41.71.105.58

# List available backups
ls -lh /opt/vuka/vuka-api.backup.*

# Stop service
sudo systemctl stop vuka-api

# Restore backup
sudo cp /opt/vuka/vuka-api.backup.20250120_123456 /opt/vuka/vuka-api
sudo chmod 750 /opt/vuka/vuka-api

# Start service
sudo systemctl start vuka-api
```

### Manual Deployment

If you need to deploy manually:

```bash
# Build locally
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o vuka-api -ldflags="-s -w" cmd/main.go
chmod 750 vuka-api

# Deploy
scp -P 2222 vuka-api yetu@41.71.105.58:/tmp/vuka-api-new
ssh -p 2222 yetu@41.71.105.58 "sudo systemctl stop vuka-api && \
  sudo mv /tmp/vuka-api-new /opt/vuka/vuka-api && \
  sudo chmod 750 /opt/vuka/vuka-api && \
  sudo chown yetu:yetu /opt/vuka/vuka-api && \
  sudo systemctl start vuka-api"
```

## Security Notes

- Binary permissions are set to `750` (owner can read/write/execute, group can read/execute)
- SSH key is securely stored in GitHub Secrets
- SSH key is cleaned up after deployment
- Connection uses custom SSH port (2222)
- Service runs as non-root user (`yetu`)

## Monitoring

After deployment, monitor your application:

```bash
# Watch service logs in real-time
ssh -p 2222 yetu@41.71.105.58 "sudo journalctl -u vuka-api -f"

# Check service status
ssh -p 2222 yetu@41.71.105.58 "sudo systemctl status vuka-api"

# View recent logs
ssh -p 2222 yetu@41.71.105.58 "sudo journalctl -u vuka-api -n 100 --no-pager"
```

## Best Practices

1. **Always test on develop branch first**
2. **Create release tags from main branch only**
3. **Use semantic versioning**
4. **Write meaningful tag messages**
5. **Monitor deployment logs**
6. **Test the API after deployment**
7. **Keep backups of working versions**
