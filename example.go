package main

import (
	"fmt"

	"github.com/7045kHz/http-client/gohttp"
)

func main() {
	client := gohttp.New()
	res, err := client.GET("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)

}
