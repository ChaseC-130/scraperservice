# Scraper Service

A simple web service in which collects and exposes some prometheus metrics to serve as a POC.

## Installation

For Local development: All source files for the go application are available under /src 
Simply go "go build -o scraperservice" under /src and run the scraperservice binary
Alternatively, you may find the Docker container:

You may need to build the binary for Linux systems from /src:
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/bin/scraperservice



### Kubernetes

You may apply the Deployment.yaml spec in which will also provide a prometheus and grafana pod 
