package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"net/http"
	"os"
)

func main() {
	var (
		url      string
		redirect bool
	)

	pflag.StringVarP(&url, "url", "u", "http://localhost/", "URL")
	pflag.BoolVarP(&redirect, "redirect", "r", false, "REDIRECT")
	pflag.Parse()

	status := getStatusCode(url)

	switch {
	case status >= 400:
		fmt.Printf("CheckHTTP CRITICAL: %d\n", status)
		os.Exit(2)
	case status >= 300 && redirect:
		fmt.Printf("CheckHTTP OK: %d\n", status)
		os.Exit(0)
	case status >= 300:
		fmt.Printf("CheckHTTP WARNING: %d\n", status)
		os.Exit(1)
	default:
		fmt.Printf("CheckHTTP OK: %d\n", status)
		os.Exit(0)
	}
}

func getStatusCode(url string) int {
	request, _ := http.NewRequest("GET", url, nil)
	response, err := http.DefaultTransport.RoundTrip(request)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return response.StatusCode
}
