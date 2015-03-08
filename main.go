package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func endpoint() string {
	endpoint := os.Getenv("HTTP_ENDPOINT")
	if endpoint == "" {
		fmt.Fprintln(os.Stderr, "The `HTTP_ENDPOINT` variable is required but was empty")
		os.Exit(1)
	}
	return endpoint
}

func main() {
	_, err := http.PostForm(endpoint(), url.Values{"msg": {"Hello"}, "from": {"dispatch"}})
	if err != nil {
		panic(err)
	}
}
