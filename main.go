package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	logC := make(chan string)
	storeC := make(chan string)
	store := []string{}

	go func() {
		for {
			store = append(store, <-storeC)
			fmt.Println(store)
		}
	}()

	go func() {
		for {
			log.Print(<-logC)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)

		d := string(body)

		storeC <- d
		logC <- fmt.Sprintf("added: %v", d)
	})

	http.ListenAndServe(":9000", nil)
}
