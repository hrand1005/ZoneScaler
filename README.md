# ZoneScaler

Distributed Game Server

![ZoneScaler Diagram](docs/ZoneScaler.png)

# Setup

Prerequisites:

-   Go Version 1.20
-   Docker

1. Clone repo:

```sh
git clone git@github.com:hrand1005/ZoneScaler.git
```

2. From repository root, run docker compose:

```sh
docker-compose up
```

3. To stop running containers:

```sh
docker-compose down
```

You might also want to build images and run containers individually:

Coordinator:

```sh
docker build -t coordinator-image -f cmd/coordinator/Dockerfile .
docker run -p 8080 coordinator-image
```

Worker:

```sh
docker build -t worker-image -f cmd/worker/Dockerfile .
docker run -p 8081 worker-image
```

Other scripts used during development are available in `scripts/`.
