package main

import (
    "log"
    "net/http"
    "dermadelight/routes"
)

func main() {
    router := routes.SetupRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
