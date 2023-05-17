package main

import (
	"io"
	"log"
	"net/http"
	"github.com/rimon-fedyuk/auth_middleware"
)

func final(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "<h1>Auth Succeed\n</h1>")
}

func main() {

    mux := http.NewServeMux()

    finalHandler := http.HandlerFunc(final)
    mux.Handle("/auth", auth_middleware.Auth(finalHandler))

    log.Print("Listening on :3000...")
    err := http.ListenAndServe(":3000", mux)
    log.Fatal(err)
}
