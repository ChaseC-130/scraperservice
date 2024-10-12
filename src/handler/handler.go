package handler

import (
    "fmt"
    "io"
    "net/http"
    "strconv"
	"encoding/json"
	"scraperservice/metrics"
	
)

type RequestBody struct {
    URL string `json:"url"`
}


func init() {
    http.HandleFunc("/", postHandler)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Non-POST request", http.StatusMethodNotAllowed)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()


	var reqBody RequestBody
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
	}

	targetUrl := reqBody.URL
	if targetUrl == "" {
		http.Error(w, "Empty URL", http.StatusBadRequest)
        return
	}
	
	client := http.Client{}
    res, err := client.Get(targetUrl)
	var responseCode string

	if err != nil {
        responseCode = "Unknown error"
        metrics.RequestCounter.WithLabelValues(targetUrl, responseCode).Inc()
        http.Error(w, fmt.Sprintf("GET error: %s", err), http.StatusInternalServerError)
        return
	}

	defer res.Body.Close()

    responseCode = strconv.Itoa(res.StatusCode)
    metrics.RequestCounter.WithLabelValues(targetUrl, responseCode).Inc()

    w.WriteHeader(http.StatusOK)
    fmt.Printf("GET: %s status code: %s\n", targetUrl, responseCode)
}
