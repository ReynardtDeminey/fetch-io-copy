package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := getURL()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}

	statusCode := resp.Status

	fmt.Printf("%v\n", statusCode)

}

func getURL() string {
	url := os.Args[1]
	if strings.HasPrefix(url, "http://") == false {
		return "http://" + url
	}
	return url
}
