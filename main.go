package main

import (
	"fmt"
	"go-framework/core/router"
	"net/http"
)

func main() {
	fmt.Println("start http server success")
	err := http.ListenAndServe(":9502", router.AutoLoadRouter())
	if err != nil {
		fmt.Println("start http server error")
	}
}
