package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/env", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(os.Getenv("SERVE_ENV_DATA")))
	})

	fmt.Printf("Starting server at: http://localhost:3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
