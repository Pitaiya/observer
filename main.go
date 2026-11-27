package main

import "net/http"

func main() {
	if err := http.ListenAndServe(":8080", http.NewServeMux()); err != nil {
		panic(err)
	}
}
