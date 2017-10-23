# Spaceport

All-in-one Vostok development infrastructure with integration tests inside.
All component images you may find at [Docker Hub](https://hub.docker.com/u/vstk/).

Spaceport is useful in two cases:

1. In combination with [launchpad](https://github.com/vostok/launchpad), to test newly created Vostok-instrumented applications.
2. To fix bugs or develop new components in Vostok.

## Get Started

[Docker](https://docs.docker.com/engine/installation/) (and `docker-compose`) are prerequisites. Check that you have them in your PATH:

```
$ docker --version
Docker version 17.09.0-ce, build afdb6d4

$ docker-compose --version
docker-compose version 1.16.1, build 6d1ac21
```

**Beware:** you won't be able to run all services on a very old or weak machine.
Mid-2014 MacBook Pro with 8GB RAM shows acceptable performance in our experience.

If you have `make`:

- `make` or `make up` will download and run all necessary containers
- `make down` will stop and remove them
- `make test` will run integration test on running containers, while `make full-test` will stop and recreate them before running tests
- `make pull` will pull latest versions of containers (and probably overwrite your local custom ones!)

If you don't have `make`, life will be a tiny bit harder for you. Just look inside the `Makefile` for commands.

## Web Applications

Spaceport provides:

- [Graphite](https://graphiteapp.org) at `localhost:6304` and [Grafana](https://grafana.com) at `localhost:6303`
- [Kibana](https://www.elastic.co/products/kibana) at `localhost:6305`
- [Contrails](https://github.com/vostok/contrails.web) at `localhost:6302`

## Test Vostok-instrumented Applications

Use [launchpad](https://github.com/vostok/launchpad) to create a boilerplate C# project. It will be preconfigured to work with Spaceport on your local machine.
Make some HTTP requests to your application and see results in Grafana, Kibana and Contrails.

## Fix Bugs and Develop New Components in Vostok

Let's say you want to add some features to Contrails. This usually requires working on [Contrails API](https://github.com/vostok/contrails.api) and [Contrails Web](https://github.com/vostok/contrails.web) at the same time.

First, you clone both Contrails repositories. `appsettings.json` in API is preconfigured to look for Kafka, Cassandra and other components by their container names. Change these names to `localhost`, since all Spaceport components are bound to the same ports on `localhost`. Now you have a working Spaceport with Contrails API replaced by your own application. Hack happily, and send us your pull requests.
