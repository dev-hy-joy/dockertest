package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received request -hy")
		fmt.Fprint(w, "Hello Docker_hy")

	})

	log.Println("start server -hy")

	http.ListenAndServe(":8080", nil)

	// server := &http.Server{
	// 	Addr:":8080"
	// }

	// if err := server.ListenAndServer();err != nil{
	// 	log.Println(err)
	// }
}
