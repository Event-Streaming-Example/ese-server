# ESE Server

This repo contains the code to run the BE server that will publish the events generated by the FE to our Kafka pipeline. It internally uses redis as a temporary data store with an expiry of 5 minutes.

Connecting redis here is meant to serve as an example of how a dummy service can leverage it alongside a streaming pipleine for their customized usecase.

A combination of Grafana and Prometheus has been used to publish the metrics of our service and observe its performance.

- __Language__ : [Go](https://go.dev/doc/)
- __HTTP Framework__ : [Gin](https://gin-gonic.com/docs/)
- __Database__ : [Redis](https://redis.io/docs/connect/clients/go/), [Prometheus (time-series)](https://prometheus.io/docs/prometheus/latest/getting_started/)
- __Observability__ : [Grafana](https://grafana.com/docs/grafana/latest/)

---

## Prerequisites

Install the below-mentioned docker images.

```bash
docker pull redis/redis-stack-server
docker pull prom/prometheus
docker pull grafana/grafana
docker pull saumyabhatt10642/ese-server
```

---

## Local Setup

In order to get the application running, just start the Redis stack and run the following commands

```bash
docker run -p 6379:6379 -d --name redis-local-server redis/redis-stack-server
docker run -it -p 2000:2000 --name ese-local-server saumyabhatt10642/ese-server
```

Or if you don't have docker and want to run the Go commands directly:

```bash
cd src
go get .
go run .
```

__Postman Collection__ : [JSON File](./files/Postman%20Collection.json)

---

## Setting up the Dashboard

Run the Grafana and the Prometheus images to set up metric collection and visualization.

```bash
docker run -p 9090:9090 --name prometheus-local-server -v ./src/properties/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
docker run -p 3000:3000 --name grafana-local-server grafana/grafana
```

Log in to the local Grafana dashboard using the following credentials:

__Username :__ admin

__Password :__ admin

Import [this](./files/ESE%20Server%20Grafana%20Dashboard.json) file to set up the dashboard. When prompted to provide a data source, connect to the
__Prometheus Datasource URL :__ `http://host.docker.internal:9090`

---

## Exposing Server to your local network

If on Linux or Mac, run the following command to get the IP of the machine on which your Gin server is running

```bash
ifconfig | grep netmask
```

- Make sure the device that the server is running and the machine you want to connect to is on the same network.
- Get the broadcast IP address that is returned as a result of the above query.
- `<ip address>:2000` is where your device could connect to access the Gin server.

---

## Pushing Image

If you have made changes to the application, publish a new image with the following commands:

[DockerHub Repository](https://hub.docker.com/repository/docker/saumyabhatt10642/ese-server/general)

```bash
docker image build -t ese-client:tag .
docker image tag ese-client:tag saumyabhatt10642/ese-client:tag
docker push saumyabhatt10642/ese-client:tag
```
