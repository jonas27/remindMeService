package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func main() {
	log.SetFlags(log.Lshortfile)
	http.HandleFunc("/store", storeHandler)
	http.HandleFunc("/retrieve", retrieveHandler)
	http.ListenAndServe(":2001", nil)
}

func storeHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		enableCors(&w)
		return
	}
	b := getBody(r)
	msg := unmarshal(b)
	user := "jonas"
	store(path.Join(pth, user+"/"+msg.Metric), msg)
}

func retrieveHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		return
	}
	b := getBody(r)
	msg := unmarshal(b)
	user := "jonas"
	w.Write([]byte(retrieveJsonObj(path.Join(pth, user+"/"+msg.Metric))))
}

func unmarshal(body []byte) Comment {
	var msg Comment
	e := json.Unmarshal(body, &msg)
	if e != nil {
		log.Println(e)
	}
	return msg
}

func getBody(r *http.Request) []byte {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	return bodyBytes
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}
