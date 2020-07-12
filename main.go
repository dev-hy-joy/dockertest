package main

import(
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
		log.Println("received request")
		fmt.Fprintln("received request")
		fmt.Fprintln(w, "Hello Docker_hy")

	})

	log.Println("start server")
	server := &http.Server{
		Addr: ":8000"
	}
	if err := server.ListenAndServer();err != nil{
		log.Println(err)
	}
}