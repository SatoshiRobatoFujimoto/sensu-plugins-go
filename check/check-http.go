package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/pflag"
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
		fmt.Println("CheckHTTP CRITICAL:", err)
		os.Exit(2)
	}

	return response.StatusCode
}
