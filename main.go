package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var checkoutRequests = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "checkout_http_requests_total",
	Help: "Checkout requests by status code.",
}, []string{"code"})

func checkout(w http.ResponseWriter, r *http.Request) {
	checkoutRequests.WithLabelValues("200").Inc()
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/checkout", checkout)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
