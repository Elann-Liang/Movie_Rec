package main

import (
    "net/http"
    //"fmt"
    "io/ioutil"
)

// The http path that returns "Hi" to user.
func hello(wr http.ResponseWriter, r  *http.Request){
    wr.Write([]byte("Hi!"))
}

// This function sends GET to google.com and prints the body in response.
func getGoodle(wr http.ResponseWriter, r  *http.Request){
    response, err := http.Get("https://www.google.com")
    if(err!=nil){
        wr.Write([]byte (err.Error()))
        return
    }
    result, err:= ioutil.ReadAll(response.Body)
    if(err!=nil){
        wr.Write([]byte (err.Error()))
        return
    }
    wr.Write([]byte(result))
}

func main() {
    // Create handler that maps funcs to paths.
    handler := http.NewServeMux()
    handler.HandleFunc("/hello", hello)
    handler.HandleFunc("/google", getGoodle)

    // Start the server.
    http.ListenAndServe(":5000", handler)
}