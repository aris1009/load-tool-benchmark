package main

import (
	"os"
	"os/signal"
	"fmt"
	"net/http"
	"math/rand"
)

var count = 0

func main() {
	c:= make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Printf("\n\nTotal Request Count: %d\n\n", count)
		os.Exit(1)
	}()

	http.HandleFunc("/", OkServer)
	http.ListenAndServe(":3005", nil)
}

func OkServer(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Println(randomNumber())
	fmt.Fprintf(w, "OK")
}

func randomNumber() float64 {
	return rand.Float64()
}