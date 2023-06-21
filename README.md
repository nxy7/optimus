![Optimus Prime](https://www.dexerto.com/cdn-cgi/image/width=3840,quality=75,format=auto/https://editors.dexerto.com/wp-content/uploads/2023/04/21/Transformers-1.jpg)

Monorepo management tool that's extensible and will fit any workflow.

# Key Features:

- Simplified monorepo management.
- Doesn't use obscure languages for configuration.
- Extensible using any shell that's available on the system.
- Discoverable: allows newcomers to easily find out all parts of the system they can interact with.
- Smart testing makes CI pipelines faster.
- Integrates with other tools like telepresence.
- Works with Docker Compose, Kubernetes and standalone apps.
- Composable: configuration can be split into many files.
- Commands work from any directory within repository.

# Why use shell to extend Optimus capabilities

I'm not huge fan of shell scripting, but when it comes to managing whole project it just makes sense. Shell scripting languages treat external executables as first class objects, which makes it easy to quickly adapt Optimus to any workflow. Most developers are familliar with at least one shell scripting language (even if only to launch their apps), so allowing them to reuse existing knowledge is much better than forcing them into using specific tech (many simmilar tools use python/starlark/lua). If you want to use sane scripting language while configuring Optimus I recommend `nushell`.

# Example configurations

## Standalone App

```yaml
# invoked via 'optimus frontend dev'
services:
  frontend:
    dev: |
      pnpm i && pnpm dev

  backend:
    dev: |
      cargo run

# invoked via 'optimus utilityFunction'
utilityFunction:
  description: |
    Function that does some repository related work
  run: |
    echo "Working..."
```

## Docker Compose

```yaml
# invoked via 'optimus start'
start: |
  docker compose -f compose.yml -f compose.dev.yml up -d

# invoked via 'optimus stop'
stop: |
  docker compose -f compose.yml -f compose.dev.yml down

# if any service contains 'build' you can use 'optimus build' to run builds in all services cocurrently
# it works the same for 'test' command
services:
  frontend:
    dev: |
      pnpm i && pnpm dev
    build: |
      docker build .

  backend:
    dev: |
      cargo run
    build: |
      docker build .
```

## Kubernetes

Here's part of configuration that I'm using for the service I'm developing that runs on Kubernetes
```yaml
# 'optimus start' invokes external script, which is useful if you have more logic to some step
start:
  description: |
    Start streampai application
  run: |
    nu ./scripts/kubernetes.nu init

clean: 
  description: |
    Delete all project resources
  run: |
    nu ./scripts/kubernetes.nu purge

telepresence-reset: 
  description: |
    Telepresence sometimes hangs and needs to be reset using this command
  run: |
    nu ./scripts/kubernetes.nu reset-telepresence     

services:
  frontend:
    dev: |
      telepresence intercept frontend --port 3000:http --mechanism tcp --namespace streampai
      pnpm i && pnpm dev
    build: |
      docker run .
  backend/main:
    root: ./backend
    dev: |
      telepresence intercept backend --port 7000:http --mechanism tcp --namespace streampai
      cargo run --bin main
    build: |
      docker run . -f Dockerfile.main
  backend/streamchat:
    root: ./backend
    dev: |
       cargo run --bin streamchat
```


# How to install

The easiest way to install this software would be using output of the flake from this repo and that's primarily how I'm using it in my other projects.

# Why 'Optimus'

For some reason I thought that if You compared microservices to transformers, then tool managing them should be called by the name of transformers leader. The name sounds familiar, is easy to remember and easy to alias ('op'). 