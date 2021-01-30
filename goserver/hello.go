package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hi there I love %s", r.URL.Path[1:])
	opsProcessed.Inc()
	headers := r.Header

	for index, element := range headers {
		fmt.Fprint(w, index)
		fmt.Fprintln(w, element)
	}
}

// func recordMetric() {
// 	go func() {
// 		for {
// 			opsProcessed.Inc()
// 			time.Sleep(2 * time.Second)
// 		}
// 	}()
// }

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total",
		Help: "The total number of requests",
	})
)

func main() {
	// recordMetric()
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)

	http.ListenAndServe(":3117", nil)

}
