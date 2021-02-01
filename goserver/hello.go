package main

import (
	"net/http"
	"sort"
	"text/template"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const portNumber = ":8080"

// the template that will be used to display the table
const doc = `
<!DOCTYPE html>
<html>
    <head>
        
    </head>
    <body>
		
		<table border="3" align="center">

		<tr>
    		<th>Key</th>
    		<th>Value</th>   
  		</tr>

			{{range .Pairs}}

				<tr>
				<td>
					{{.Key}}		
					{{range .Value}}
						<td>
							{{.}}
						</td>
					{{end}}
				</td>
				</tr>
			{{end}}
		</table>
       
    </body>
</html>`

// This is to store each key value pair from the headers
type Pair struct {
	Key   string
	Value []string
}

//This is to store all the Pairs from "Pair" struct
type Keys struct {
	Pairs []Pair
}

func handler(w http.ResponseWriter, r *http.Request) {

	opsProcessed.Inc() //increase counter for metric total by 1

	headers := r.Header
	w.Header().Add("Content Type", "text/html")
	templates := template.New("template")
	templates.New("doc").Parse(doc)

	pairs := Keys{}

	//sorting the keys first
	var sortedKey []string

	for index, _ := range headers {

		sortedKey = append(sortedKey, index)

	}
	sort.Strings(sortedKey)

	// using sorted keys to access each element from request headers, and store them in the struct
	for _, element := range sortedKey {
		pair := Pair{}
		pair.Key = element

		for _, value := range headers[element] {
			pair.Value = append(pair.Value, value)
		}

		pairs.Pairs = append(pairs.Pairs, pair)
	}

	//generate the template to be displayed
	templates.Lookup("doc").Execute(w, pairs)

}

// for prometheus custom metric
var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total",
		Help: "The total number of requests",
	})
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)
	http.ListenAndServe(portNumber, nil)

}
