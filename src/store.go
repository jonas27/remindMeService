package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"time"
)

var (
	pth = "/volumes/remindMeService/user/"
)

type Comment struct {
	Comment string    `json:"comment,omitempty"`
	Data    float64   `json:"data"`
	Metric  string    `json:"metric"`
	UserID  string    `json:"userID"`
	Time    time.Time `json:"time,omitempty"`
}

func (c *Comment) print() string {
	return "me"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func store(pth string, c Comment) bool {
	log.Println("store: " + pth)
	f := openWrite(pth)
	c.Time = time.Now()
	defer f.Close()
	j, e := json.Marshal(c)
	if e != nil {
		log.Println(e)
		return false
	}
	str := string(j) + "\n"
	_, e = f.WriteString(str)
	if e != nil {
		log.Println(e)
	}
	return true
}

func openWrite(pth string) *os.File {
	e := os.MkdirAll(path.Dir(pth), os.ModePerm)
	if e != nil {
		log.Println(e)
	}
	f, e := os.OpenFile(pth, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if e != nil {
		log.Println(e)
	}
	return f
}
