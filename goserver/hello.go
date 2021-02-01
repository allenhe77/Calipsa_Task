package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const doc = `
<!DOCTYPE html>
<html>
    <head>
        
    </head>
    <body>
		
		<table>
			{{range .Keys}}

			<tr>
				<td>{{.}}</td>

				
			</tr>

			
			
		{{end}}
		</table>
       
    </body>
</html>`

type Pair struct {
	Key   string
	Value []string
}

type Keys struct {
	Pairs []Pair
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hi there I love %s", r.URL.Path[1:])
	opsProcessed.Inc()
	headers := r.Header
	// w.Header().Add("Content Type", "text/html")
	// templates := template.New("template")
	// templates.New("doc").Parse(doc)
	a := Keys{}

	for index, element := range headers {
		// fmt.Fprint(w, reflect.TypeOf(index))
		// fmt.Fprintln(w, reflect.TypeOf(element))

		pair := Pair{}
		pair.Key = index

		for _, value := range element {
			pair.Value = append(pair.Value, value)
		}
		// for c, d := range element {

		// }
		a.Pairs = append(a.Pairs, pair)

	}

	fmt.Println(a)

	// templates.Lookup("doc").Execute(w, a)

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
	// http.HandleFunc("/", handler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
