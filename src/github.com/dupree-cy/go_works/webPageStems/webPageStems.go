package main

import (
        "fmt"
        "net/http" 
        "io/ioutil"
       )

func main() {
  url      := "http://revolutionanalytics.com/why-revolution-analytics"
  res, err := http.Get(url)
  if err != nil {
   fmt.Println("http.Get", err)
   return
  }
  defer res.Body.Close()
  fmt.Println(url)
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println("ioutil.ReadAll", err)
    return
  }
  fmt.Println(string(body))
} 
