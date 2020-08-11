# Go SDK to Prometheus to Cortex Demo

This demo sends data from a Go program to Prometheus, which then writes to Cortex using
the `remote_write` API.

## Setup
### Prometheus
An instance of Prometheus must be running at port 8888 in order for Prometheus to receive
metrics data from the Go application. The demo works with the provided sample
`prometheus.yml` file.

Install instructions:
```shell
git clone https://github.com/prometheus/prometheus
cd prometheus
go build ./cmd/prometheus
./prometheus --config.file=prometheus.yml
```

### Cortex
An instance of Cortex must be running to receive data from the `remote_write` API. 

Install Instructions:
```shell
git clone git@github.com:cortexproject/cortex.git
go build ./cmd/cortex
./cortex -config.file=./docs/configuration/single-process-config.yaml
```

### Grafana
The demo verifies that the export and remote_write succeeded by checking the data on
Grafana.

Install Instructions:
```shell
docker run --rm -d --name=grafana -p 3000:3000 grafana/grafana
```

## Checking Values
1. Go to `localhost:3000`
2. Login with username: admin and password: admin (skip password change)
3. Add the `remote_write` URL (`http://host.docker.internal:9009/api/prom`) as a data
   source. 
4. Add a dashboard and query either `a_counter` or `a_value_recorder` to see the results.