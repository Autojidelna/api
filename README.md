# CORE example api

## Running the project

1. Download [golang](https://go.dev/dl/)

2. Add go bin to your path

for macos/linux

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc
source ~/.bashrc
source ~/.zshrc
```

for windows

```cmd
setx PATH "%PATH%;C:\Go\bin"
setx PATH "%PATH%;%USERPROFILE%\go\bin"
```

```bash
go install github.com/air-verse/air@latest
```

After that just run the project with

```bash
air
```

## Developer tools

- [ent](https://entgo.io/docs/getting-started)
  - for generating models. To regenerate models run `go generate ./ent`
- [air](https://github.com/air-verse/air)
  - for hot reloading the server. To run the server run `air`
- [swag](https://github.com/swaggo/gin-swagger)
  - for generating swagger documentation. To regenerate swagger documentation run `swag init`
  - install with `go install github.com/swaggo/swag/cmd/swag@latest`

### vscode extensions

- [golang](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [Run on Save](https://marketplace.visualstudio.com/items?itemName=emeraldwalk.RunOnSave)
  - We use this to automatically generate ent models and swagger documentation on save

## Deployment

1. Download [docker](https://docs.docker.com/get-docker/)
2. Enable docker buildx - just check the box `Use containerd for pulling and storing images` in docker desktop -> settings -> general
3. Build the image

```bash
docker buildx build --platform linux/amd64,linux/arm64 -t registry.appelevate.cz/coree:latest .
```

4. Push the image to the registry

```bash
docker push registry.appelevate.cz/coree:latest
```

5. Deploy the image to the server

```bash
ssh ohp
docker compose up -d
```

## Setting up the alpine server

1. Install docker

```bash
apk add docker
```

2. Install docker-compose

```bash
apk add docker-compose
```

3. Run cloudflare tunnel

```bash
docker run -d \
  --restart always \
  --network bridge \
  cloudflare/cloudflared:latest tunnel --no-autoupdate run --token <your-token-here>
```

4. Run the registry

```bash
mkdir registry
cd registry
mkdir data
mkdir auth
vim docker-compose.yml
```

```yaml
version: "3"
services:
  registry:
    image: registry:2
    ports:
      - "5000:5000"
    environment:
      REGISTRY_AUTH: htpasswd
      REGISTRY_AUTH_HTPASSWD_REALM: Registry
      REGISTRY_AUTH_HTPASSWD_PATH: /auth/registry.password
      REGISTRY_STORAGE_FILESYSTEM_ROOTDIRECTORY: /data
    volumes:
      - ./auth:/auth
      - ./data:/data
    restart: always
```

```bash
apk add apache2-utils
cd auth
htpasswd -Bc registry.password username
cd ..
docker-compose up -d
cd /root/
vim deploy_docker.sh
```

```bash
cd /root/clone/*/
docker compose down
cd /root/
rm -rf clone
mkdir clone
cd clone
git clone --depth 1 git@github.com:app-elevate/CORE.go-gin-api.git
cd */
docker compose up -d
```

Add the github ssh key to the server

```bash
vim /root/.ssh/id_rsa
```

Add the trusted ssh client

```bash
ssh-keyscan github.com >> /root/.ssh/known_hosts
```

add the authorized key to the deployment

```bash
vim /root/.ssh/authroized_keys
```

first is the key you use to access ssh. The second one is the key that is used to deploy the docker

```
ssh-ed25519 AAAAKEY
no-agent-forwarding,no-X11-forwarding,no-port-forwarding,no-pty,command="/root/deploy_docker.sh" ssh-ed25519 AAAAKEY
```
