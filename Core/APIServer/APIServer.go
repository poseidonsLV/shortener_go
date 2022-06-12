package APIServer

import (
	"log"
	"net/http"
)

func Run() {
	startWebServer()
}

func startWebServer() {
	log.Fatal(http.ListenAndServe(":8001", nil))
}
