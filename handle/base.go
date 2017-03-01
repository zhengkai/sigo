package handle

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/zhengkai/sigo/layout"
)

type BaseHandle struct {
	uri    string
	layout layout.Layout
	data   interface{}
}

func (BaseHandle) DefaultLayout() layout.Layout {
	return &layout.BaseLayout{}
}

func (this *BaseHandle) SetLayout(l layout.Layout) {
	this.layout = l
}

func (this *BaseHandle) SetUri(s string) {
	this.uri = s
	fmt.Println(`set uri =`, this.uri)
}

func (this *BaseHandle) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	this.data = make(map[string]interface{})
}

func (this *BaseHandle) Parse() *bytes.Buffer {
	if this.layout == nil {
		this.layout = this.DefaultLayout()
	}
	this.layout.SetPath(this.uri)
	return this.layout.Parse()
}
