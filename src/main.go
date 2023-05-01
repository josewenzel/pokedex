package main

import (
	"fmt"
	"log"
	"net/http"
	"pokedex/src/api/handler"
	"strconv"
)

func main() {
	port := 1234
	http.HandleFunc("/status", handler.StatusHandler)
	fmt.Printf("Running in http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
