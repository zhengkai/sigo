package handle

import (
	"bytes"
	"net/http"

	"github.com/zhengkai/sigo/layout"
)

type Handle interface {
	SetUri(string)
	Get(w http.ResponseWriter, r *http.Request)
	Parse() *bytes.Buffer
	SetLayout(layout.Layout)
	DefaultLayout() layout.Layout
}

func Register(uri string, data Handle) {
	http.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		d := data
		d.SetUri(uri)
		d.Get(w, r)
		w.Write(d.Parse().Bytes())
	})
}
