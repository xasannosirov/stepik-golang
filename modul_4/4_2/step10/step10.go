package main

import (
    "fmt"
	"io"
	"net/http"
	"net/url"
	// "time"
    "log"
    "os"
)

func main() {
    var name,age string
    fmt.Scan(&name, &age)
    baseURL := "http://127.0.0.1:8080/hello"
    params:=url.Values{}
    params.Add("name", name)
    params.Add("age", age)
    fullURL:=baseURL+"?"+params.Encode()
    
    response, err:=http.Get(fullURL)
    if err!=nil{
        log.Fatalln(err)
        return
    }
    defer response.Body.Close()
    _,err=io.Copy(os.Stdout, response.Body)
    if err!=nil{
        log.Fatalln(err)
    }
}