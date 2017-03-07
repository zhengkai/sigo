package handle

import (
	"bytes"
	//"fmt"
	"net/http"
	"reflect"

	"github.com/zhengkai/sigo/layout"
)

type Handle interface {
	SetUri(string)
	Get(w http.ResponseWriter, r *http.Request)
	Parse() *bytes.Buffer
	SetLayout(layout.Layout)
	DefaultLayout() layout.Layout
	Clone() Handle
}

func Register(uri string, data Handle) {

	// fmt.Println(c, data)
	// fmt.Printf("%T %T\n", c, data)

	http.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {

		d := reflect.New(reflect.ValueOf(data).Elem().Type()).Interface().(Handle)
		// fmt.Printf("%T %T\n", d, c)
		d.SetUri(uri)
		d.Get(w, r)
		w.Write(d.Parse().Bytes())
	})
}
