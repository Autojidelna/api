# autojidelna example api

## Running the project

1. Download [golang](https://go.dev/dl/)

2. Download and run [postgres](https://postgresapp.com/downloads.html)

3. Add go bin to your path

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

after that install air with:

```bash
go install github.com/air-verse/air@latest
```

create a temp directory for go executables

```bash
sudo mkdir -p /usr/local/go/bin/tmp
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
docker buildx build --platform linux/amd64,linux/arm64 -t registry.autojidelna.cz/autojidelna:latest .
```

4. Push the image to the registry

```bash
docker push registry.autojidelna.cz/autojidelna:latest
```

5. Deploy the image to the server

```bash
ssh autojidelna # ssh onto the server if you have access
sh deploy_docker.sh
```
