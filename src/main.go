package main

import (
        "net/http"
        "log"
        "os"
        "fmt"
        _ "scraperservice/handler"
        _ "scraperservice/metrics"
)

func main() {

    apiport := os.Getenv("API_PORT")
    if apiport == "" {
        apiport = "8080"
    }

    metricsport := os.Getenv("METRICS_PORT")
    if metricsport == "" {
        metricsport = "9095"
    }

    // API server
    go func() {
        fmt.Println("API server is listening on port: ", apiport)
        err := http.ListenAndServe(":" + apiport, nil)
        if err != nil {
            log.Fatalf("Error starting api server")
        }
    }()


    // Metrics server
    fmt.Println("Metrics server is listening on port: ", metricsport)
    err := http.ListenAndServe(":" + metricsport, nil)
    if err != nil {
        log.Fatalf("Error starting metrics server")
    }
}