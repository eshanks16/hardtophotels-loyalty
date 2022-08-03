package main

import (
	"fmt"
	"log"

	//html templates
	"net/http"

	//prometheus
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func main() {

	//web
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.Handle("/metrics", promhttp.Handler())
	
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
