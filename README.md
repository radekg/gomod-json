# gomod-json

Read go.mod file contents and dump them as JSON.

## Install

```sh
make install
```

## Why

For example, tagging your Docker images based on some dependency:

```sh
git clone https://github.com/radekg/jsonnet-playground.git .
JSONNET_VERSION=(gomod-json --path=./go.mod | yq '.Require[] | select(.Mod.Path == "github.com/google/go-jsonnet") | .Mod.Version')
docker build -t docker.io/radekg/jsonnet-playground:${JSONNET_VERSION}
docker tag docker.io/radekg/jsonnet-playground:${JSONNET_VERSION} docker.io/radekg/jsonnet-playground:latest
```
