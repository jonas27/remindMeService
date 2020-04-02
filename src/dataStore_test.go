package main

import (
	"bufio"
	"log"
	"os"
	"testing"
	"time"
)

func logLines() {
	log.SetFlags(log.Lshortfile)
}

func TestStore(t *testing.T) {
	logLines()
	c := Comment{
		Metric:  "test",
		Data:    12,
		UserID:  "test",
		Comment: "test",
		Time:    time.Now(),
	}
	path := "/volumes/remindMeService/user/id1/series1/comments"
	if !store(path, c) {
		log.Fatal()
	}
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}

}
