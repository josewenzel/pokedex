package main

import (
	"fmt"
	"log"
	"net/http"
	"pokedex/src/api/router"
	"strconv"
)

func main() {
	port := 1234
	server := http.NewServeMux()

	SubscribeStatusHandler(server)

	fmt.Printf("Running in http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func SubscribeStatusHandler(server *http.ServeMux) {

	server.HandleFunc("/status", router.RouteStatusRequest)
}
