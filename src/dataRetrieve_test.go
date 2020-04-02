package main

import (
	"log"
	"testing"
)

func TestRetrieve(t *testing.T) {
	logLines()
	pth := "/volumes/remindMeService/user/id1/series1/comments"
	s := retrieveJsonObj(pth)
	log.Println(s)

}
