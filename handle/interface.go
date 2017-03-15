package handle

import (
	"bytes"
	// "fmt"
	"net/http"
	"time"
	// "reflect"

	"github.com/zhengkai/sigo/layout"
)

type Handle interface {
	SetUri(string)
	SetStartTime(time.Time)
	Parse()
	Output() *bytes.Buffer
	SetHttp(w http.ResponseWriter, r *http.Request)
	Prepare() bool
	New() Handle
	SetLayout(layout.Layout)
	DefaultLayout() layout.Layout
}

type Head interface {
	New() Head
	AddJS(string)
	AddCSS(string)
	AddMeta(string)
	Export() map[string]interface{}
}

func Register(uri string, h Handle) {

	// fmt.Println(x)

	http.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {

		t := time.Now()

		// d := reflect.New(reflect.ValueOf(h).Elem().Type()).Interface().(Handle)

		d := h.New()
		// fmt.Println(d)

		d.SetUri(uri)
		d.SetHttp(w, r)
		if !d.Prepare() {
			return
		}
		d.SetStartTime(t)
		d.Parse()
		w.Write(d.Output().Bytes())
	})
}
