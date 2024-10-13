# Scraper Service

A simple web service in which collects and exposes some prometheus metrics to serve as a POC.

## Installation

You may find the docker container for the scraperservice: [chasec130/scraperservice:latest](https://hub.docker.com/repository/docker/chasec130/scraperservice/general)


If you would like to compile the code yourself, all required files are in `/src`, you may simply run "go build"

The provided dockerfile under /build uses Alpine, therefor you may need to compile as follows from ARM mac:
```
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ../build/bin/scraperservice
```

The scraper service requires port 8080 for the API and 9095 for /metrics endpoint:
```
docker run -p 8080:8080 -p 9095:9095 scraperservice
```


## Deployment

In order to leverage the Prometheus configuration, you can simply apply the deployment, service, and configmap specs to your respective kubernetes cluster. The specs are presently setup in the default namespace.

This was tested in a Minikube environment, the service is a LoadBalancer thus will require `minikube tunnel` when using a minikube setup in order to route traffic correctly


## PromQL

The kubernetes deployment includes a basic promtheus container to run in the same scrapersvc pod. The webui is listening on port 9090. 

You may run an example PromQL query such as:
`http_get{code="200"}` in order to see the total count of http_get requests that returned status code `200` 

Some other examples:

`http_get{url=~".*google.com"}` to regex match any requests passed to google.com domain

`rate(http_get{code="200"}[5m])` to show the rate of how many requests per second over the last 5 minutes returned 200s

`increase(http_get{code="200"}[10m])` to show the sum of requests over the last 10 minutes that returned 200s 
Note: Due to interpolation, you may want to round the result `round(increase(http_get{code="200"}[10m]))`

![Example](/resources/promscreenshot.png)


## Scripts

query.py will query a random URL from the predefined list every 5 seconds when ran. One of the URLs is invalid to showcase 5xx responses.

