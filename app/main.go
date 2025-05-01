package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    log.Println("Servidor rodando em http://localhost:5500")
    err := http.ListenAndServe(":5500", nil)
    if err != nil {
        log.Fatal(err)
    }
}
