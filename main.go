package main

import (
	"fmt"
	"ltv_challenge/routes"
	"net/http"
)

const PORT = "5000"
const IP = "127.0.0.1"

func main() {
	mux := routes.Server_init()
	fmt.Println("Starting server")
	http.ListenAndServe(IP+":"+PORT, mux)
}
