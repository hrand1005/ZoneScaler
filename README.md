# ZoneScaler: Distributed Game Server Framework

## Overview

ZoneScaler is a robust framework tailored for the development of partitionable, scalable game servers. It facilitates the creation of distributed game server architectures where game spaces are dynamically partitioned, and server load is intelligently balanced.

## Architecture

![ZoneScaler Diagram](docs/ZoneScaler.png)

ZoneScaler operates on a coordinator/worker model:

- **Coordinator Node**: Orchestrates the server landscape, managing game partitions, and distributing loads. It is responsible for the dynamic allocation of partitions across available worker servers based on current load metrics.

- **Worker Servers**: Independently manage state and handle game logic for assigned partitions. They communicate with proxied entities in adjacent partitions for cross-partition player interactions.

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

# Configuration

Example config files for each service are found in `cmd/<service_name>/config.json`.
These typically include server configurations such as host and port numbers.
Some services provide a simple frontend for visualization, diagnostics, logs etc.
In these caess, the root directory for 'statically served' files is set by
the `static_dir` field. For an example, check out [cmd/coordinator/config.json](cmd/coordinator/config.json).

## Load Balancing and Fault Tolerance

- **Dynamic Load Balancing**: The coordinator actively monitors load across partitions and redistributes them to maintain optimal performance and balance.

- **State Preservation**: ZoneScaler employs strategies to ensure fault tolerance, such as backup servers mirroring game state or stateless server pools that can quickly take over in the event of a server failure.

## Cross-Partition Interaction

- **Proxied Entities**: Interactions across partition boundaries are facilitated through proxied entities, minimizing latency and offloading the coordinator.

- **gRPC Communication**: Entities within different partitions communicate using gRPC, ensuring efficient and direct state synchronization.

## Integration and Usage

ZoneScaler can be utilized as a Go library for building game servers or as a standalone service interfacing via gRPC. This dual-mode operation enables integration with game servers written in various programming languages.

## Community and Contributions

ZoneScaler is open-source and thrives on community contributions. Whether you are contributing code, documentation, or feedback, your involvement is crucial to its success.

- **Code**: Enhance the framework with new features or optimizations.
- **Documentation**: Improve the getting-started experience for new users by refining the existing documentation.
- **Testing**: Help test new releases and provide critical feedback.

For contribution opportunities, please see the [issues section](#).

## Roadmap

Our roadmap is focused on expanding ZoneScaler's capabilities, including but not limited to:

- Advanced partitioning analytics for better load predictions.
- More granular partition control.
- Enhanced recovery strategies for state management.

Stay connected for future developments as we continue to advance the state of distributed gaming servers.