package main

import (
    "fmt"
    "log"
    "net/http"
	"time"
)

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    }

func homePage(w http.ResponseWriter, r *http.Request){
        enableCors(&w)
	var ttt = time.Now()
    zone, offset := ttt.Zone()
    fmt.Println(zone, offset)
    fmt.Fprintln(w, "server epoch time is as :", ttt)
    fmt.Fprintln(w, "server time zone is as :", zone)
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}