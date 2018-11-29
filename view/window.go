package view

import (
	"github.com/zserge/webview"

	"bytes"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"
)

func Init(sets *webview.Settings) {
	url := startServer()
	sets.URL = url
	startWebview(sets)
}

func startWebview(sets *webview.Settings) {
	w := webview.New(*sets)
	defer w.Exit()
	w.Run()
}

func startServer() string {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if len(path) > 0 && path[0] == '/' {
				path = path[1:]
			}
			if path == "" {
				path = "index.html"
			}
			if bs, err := Asset("assets/" + path); err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
				io.Copy(w, bytes.NewBuffer(bs))
			}
		})
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}
