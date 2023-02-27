package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	proxy := NewGetter("code.whyyou.me", "git", "http://xian.whyyou.me:3000")
	server := &http.Server{
		Addr:         ":80",
		Handler:      proxy,
		ReadTimeout:  1 * time.Hour,
		WriteTimeout: 1 * time.Hour,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("启动")
}
