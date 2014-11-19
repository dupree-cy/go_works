package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Only to print the content of the page ; see goquery
func main() {
	url := "http://revolutionanalytics.com/why-revolution-analytics"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	fmt.Println(url)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll", err)
		return
	}
	fmt.Println(string(body))
	return
}
