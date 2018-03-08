package main

import (
	"net/http"
	"log"
	"goBook/ch1/exercises/lissaj"
	"strconv"
)
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	queryz := r.URL.Query()
	temp, _ := strconv.Atoi(queryz.Get("cycles"))

	lissaj.LissajousF(w, temp)
}
