package main

import (
	"fmt"
	"go-atlas-corp/server"
	"net/http"
)

func main() {
	fmt.Println("Server starting...")
	newServer := server.Server{}
	http.ListenAndServe(fmt.Sprintf(":%d", 3000), server.NewRouter(&newServer))
}
