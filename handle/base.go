package handle

import (
	"bytes"
	// "fmt"
	"net/http"

	"github.com/zhengkai/sigo/layout"
)

type BaseHandle struct {
	uri    string
	layout layout.Layout
	Data   interface{}
}

func (BaseHandle) DefaultLayout() layout.Layout {
	return &layout.BaseLayout{}
}

func (this *BaseHandle) SetLayout(l layout.Layout) {
	this.layout = l
}

func (this BaseHandle) Clone() Handle {
	c := this
	return &c
}

func (this *BaseHandle) SetUri(s string) {
	this.uri = s
	// fmt.Println(`set uri =`, this.uri)
}

func (this *BaseHandle) Get(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL.Query())
	this.Data = make(map[string]interface{})
}

func (this *BaseHandle) Parse() *bytes.Buffer {
	if this.layout == nil {
		// fmt.Println(`new Parse`)
		this.layout = this.DefaultLayout()
	}
	this.layout.SetPath(this.uri)
	this.layout.SetData(this.Data)
	return this.layout.Parse()
}
