# Coinwallet

A monolithic server that helps synchronize Blockchain transactions for chosen addresses using blockchain.com's Data API.

## Tech Stack

- [Pocketbase](https://pocketbase.io/) - web server with CRUD defaults and an intuitive UI for Admin panel to define collections, Auth, etc.
- [Go](https://go.dev/) - The accompanying code /sync API is written in Go, which plugs into the routing mechanisms provided by Pocketbase
- [Docker](https://www.docker.com/products/docker-desktop/) - Docker image is used for isolating dependencies and execution run time

## Getting Started

- Install Docker and `make start`
- Alternatively, if dependencies are already installed, use `make local`

## Common Usecases

- Go to [Admin](http://127.0.0.1:8090/_/) panel
  - Create Admin user
  - Create basic collections
  - Perform other activities like monitoring, auth, etc.
- Use `curl` for Syncing a known `address_id`
  - `http://127.0.0.1:8090/api/collections/addresses/records/z2s6gwle0laatxv/sync`
