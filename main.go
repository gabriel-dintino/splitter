package main

import (
	"fmt"
	"net/http"
	"runtime"

	configurations "configurations"
	v1 "routers"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/spliter/v1/ping", v1.PingV1)
	http.HandleFunc("/spliter/v1/ingest", v1.IngestFileV1)
	fmt.Println(fmt.Sprintf("Server started at: %s", configurations.Port))
	http.ListenAndServe(fmt.Sprintf(":%s", configurations.Port), nil)
}
