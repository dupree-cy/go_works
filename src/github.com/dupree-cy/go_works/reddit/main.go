package main

import (
    //"io"
    "log"
    "net/http"
    //"os"
    "encoding/json"
    "fmt"
)

type Item struct {
    Title string
    URL   string
}

type Response struct {
    Data struct {
        Children []struct {
            Data Item
        }
    }
}

func main() {
  resp, err := http.Get("http://reddit.com/r/golang.json") // go function can return multiple arguments
  if err != nil {
    log.Fatal(err)
  }

  if resp.StatusCode != http.StatusOK {
    log.Fatal(resp.Status)
  }

  // _, err = io.Copy(os.Stdout, resp.Body) // not so comprehensible
  r   := new(Response)
  err  = json.NewDecoder(resp.Body).Decode(r)
  for _, child := range r.Data.Children {  // print the data
    fmt.Println(child.Data.Title)
  }
  if err != nil {
    log.Fatal(err)
  }
}
