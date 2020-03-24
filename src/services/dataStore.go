package services

import (
	"fmt"
	"time"
)

type Comment struct {
	name      string
	data      int
	timestemp time.Time
}

// func Store(c Comment) {
// 	fmt.Println(c)
// }

func need() {
	fmt.Print("hallo from test")
}
