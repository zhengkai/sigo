package handle

import (
	"bytes"
	"time"
	//"fmt"
	"net/http"
	"reflect"

	"github.com/zhengkai/sigo/layout"
)

type Handle interface {
	SetUri(string)
	SetStartTime(time.Time)
	Get(w http.ResponseWriter, r *http.Request)
	Parse() *bytes.Buffer
	SetLayout(layout.Layout)
	DefaultLayout() layout.Layout
}

func Register(uri string, h Handle) {

	http.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {

		t := time.Now()

		d := reflect.New(reflect.ValueOf(h).Elem().Type()).Interface().(Handle)

		d.SetUri(uri)
		d.SetStartTime(t)
		d.Get(w, r)
		w.Write(d.Parse().Bytes())
	})
}
