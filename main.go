package main

import "net/http"

func main() {
	http.ListenAndServe(":9502", nil)
}
