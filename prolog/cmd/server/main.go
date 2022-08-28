package main

import (
	"log"
	"fmt"
	"github.com/smg061/prolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	fmt.Println("Now serving in 8080!")
	log.Fatal(srv.ListenAndServe())
}