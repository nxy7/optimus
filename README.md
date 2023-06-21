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
- Composable: configuration can be split into many files

# Why use shell to extend Optimus capabilities

I'm not huge fan of shell scripting, but when it comes to managing whole project it just makes sense. Shell scripting languages treat external executables as first class objects, which makes it easy to quickly adapt Optimus to any workflow. Most developers are familliar with at least one shell scripting language (even if only to launch their apps), so allowing them to reuse existing knowledge is much better than forcing them into using specific tech (many simmilar tools use python/starlark/lua). If you want to use sane scripting language while configuring Optimus I recommend `nushell`.

# Example configurations

## Standalone App

## Docker Compose

## Kubernetes

# How to install

The easiest way to install this software would be using output of the flake from this repo and that's primarily how I'm using it in my other projects.

# Why 'Optimus'

For some reason I thought that if You compared microservices to transformers, then tool managing them should be called by the name of transformers leader. The name sounds familiar, is easy to remember and easy to alias ('op'). 