package main

import (
	"github.com/molizz/webview-bootstrap/view"
	"github.com/zserge/webview"
)

func main() {
	view.Init(&webview.Settings{
		Title:  "",
		Width:  600,
		Height: 500,
		Debug:  true,
		ExternalInvokeCallback: func(w webview.WebView, data string) {

		},
	})
}
