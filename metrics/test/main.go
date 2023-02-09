package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			for {
				now := time.Now()
				http.Get("http://localhost:28888/hello")
				fmt.Println(time.Since(now))
				time.Sleep(1000 * time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Hour)

}
