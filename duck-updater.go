package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {

	domain := flag.String("domain", "", "Duck DNS domain")
	token := flag.String("token", "", "Duck DNS token")

	flag.Parse()

	if *domain == "" || *token == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var buffer bytes.Buffer
	buffer.WriteString("https://www.duckdns.org/update?domains=")
	buffer.WriteString(*domain)
	buffer.WriteString("&token=")
	buffer.WriteString(*token)

	resp, err := http.Get(buffer.String())

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	if string(body) == "OK" {
		fmt.Printf("Status: %s\n", strconv.Itoa(resp.StatusCode))
		fmt.Printf("Success! Updated Domains:\n")
		fmt.Printf("1. %s.duckdns.org\n", *domain)
	} else {
		fmt.Printf("Status: %s\n", strconv.Itoa(resp.StatusCode))
		fmt.Printf("Failure! Failed to update %s\n", *domain)
	}

}
