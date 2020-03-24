package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	// services "./services/dataStore.go"
	"github.com/jonas27/remindMeService/src/services"
)

func main() {
	now := time.Now()
	_ = services.Comment{}
	log.Println(now)
	// store.dataStore.need()
	// i := store
	http.HandleFunc("/", fooHandler)
	http.ListenAndServe(":3000", nil)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	// const add = me + strconv.Itoa(num)
	fmt.Fprintf(w, "Scrape metrics under \"<port>/metrics\". \n This is version 1")
}
