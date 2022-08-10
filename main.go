package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHAndler)
	log.Println("web port 9000")
	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)

}

func helloHAndler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello sir"))
}
