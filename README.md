# ESE Server

This repo contains the code to run the BE server that will publish the events generated by the FE to our Kafka pipeline. It internally uses Redis as a temporary data store with an expiry of 5 minutes.

Connecting Redis here is meant to serve as an example of how a dummy service can leverage it alongside a streaming pipeline for their customized use-case.

A combination of Grafana and Prometheus has been used to publish the metrics of our service and observe its performance.

- **Language** : [Go](https://go.dev/doc/)
- **HTTP Framework** : [Gin](https://gin-gonic.com/docs/)
- **Database** : [Redis](https://redis.io/docs/connect/clients/go/), [Prometheus (time-series)](https://prometheus.io/docs/prometheus/latest/getting_started/)
- **Observability** : [Grafana](https://grafana.com/docs/grafana/latest/)

---

## Running Locally

Before running the below commands, uncomment the Redis configs in the [configs.yaml](./src/properties/config.yaml) so that your Redis URL is pointing to where it is running locally on your machine.

```bash
# This will start Redis locally on Mac on localhost:6379. Make sure to change your settings accordingly in the config.yaml file
brew services start redis

# Move to the working folder
cd src

# Install the required dependencies
go get .

# Specify the port on which you want the application to run
SERVER_PORT=2000 go run .
```

**Note:** To get the metrics, start your Prometheus and Grafana servers separately.

**Postman Collection** : [JSON File](./files/Postman%20Collection.json)

---

## Running via Docker

If one doesn't want to run the below-given commands and directly use docker-compose: `docker-compose -p ese-sever-singular up -d`.

```bash
# First create a network where all the containers will reside in
docker network create ese-server-network

# Start the Redis container and attach it to our custom network
docker run -d --name redis-server --network ese-server-network -p 6379:6379 Redis

# Start the ESE Server on port 2000 and attach it to our custom network
docker run -d --name ese-server1 --network ese-server-network -e SERVER_PORT=2001 -p 2001:2001 saumyabhatt10642/ese-server

# Start the Prometheus Server and attach it to our custom network
docker run -d --name prometheus-server --network ese-server-network -p 9090:9090 -v $(pwd)/src/properties/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus

# Start the Grafana Server and attatch it to our custom network
docker run -d --name grafana-server --network ese-server-network -p 3000:3000 grafana/grafana
```

---

## Connecting to Grafana

Log in to the local Grafana dashboard using the following credentials. The dashboard is hosted at `http://localhost:3000`

**Username :** admin

**Password :** admin

Import [this](./files/ESE%20Server%20Grafana%20Dashboard.json) file to set up the dashboard. When prompted to provide a data source, connect to the
**Prometheus Datasource URL :** `http://prometheus-server:9090`

---

## Pushing Image

On merging any PR to master will automatically trigger a push to the `latest` tag. If however any other image tag is to be published, use the below given commands:

[DockerHub Repository](https://hub.docker.com/repository/docker/saumyabhatt10642/ese-server/general)

```bash
docker image build -t ese-server:tag .
docker image tag ese-server:tag saumyabhatt10642/ese-server:tag
docker push saumyabhatt10642/ese-server:tag
```
