package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Getter struct {
	host string //gitlab address
	vcs  string //git
	root string //git
}

func NewGetter(host string, vcs string, root string) *Getter {
	return &Getter{
		host: host,
		vcs:  vcs,
		root: root,
	}
}

func (x *Getter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("go-get") == "1" {
		sp := strings.Split(r.URL.Path[1:], "/")
		if len(sp) < 2 {
			http.Error(w, fmt.Errorf("unsupport path: %s", r.URL.Path).Error(), http.StatusBadRequest)
			return
		}
		prefix := fmt.Sprintf("%s/%s/%s", x.host, sp[0], sp[1])
		repository := fmt.Sprintf("%s/%s/%s.%s", x.root, sp[0], sp[1], x.vcs)
		fmt.Fprintf(w, `<html><head><meta name="go-import" content="%s %s %s" /></head></html>`, prefix, x.vcs, repository)
		log.Println("go get [", prefix, "] from repository [", repository, "].")
		return
	}
	http.Error(w, fmt.Errorf("unsupport request: %s", r.URL.Path).Error(), http.StatusBadRequest)
}
