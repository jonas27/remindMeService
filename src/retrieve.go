package main

import (
	"bufio"
	"log"
	"os"
)

func retrieveJsonObj(pth string) string {
	log.Println("retrieve " + pth)
	f, _ := os.Open(pth)
	scanner := bufio.NewScanner(f)
	s := ""
	for scanner.Scan() {
		s = s + scanner.Text() + ","
	}
	if len(s) == 0 {
		return "[ ]"
	}
	s = s[:len(s)-1]
	s = "[ " + s + " ]"
	return s
}
