package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	e := &ebook{
		name:  "The little fox",
		date:  time.Now(),
		study: false,
	}
	err := CreateEbook(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created! :)")
}
