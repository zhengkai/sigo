package handle

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/zhengkai/sigo/layout"
)

type BaseHandle struct {
	Time     time.Time
	Uri      string
	layout   layout.Layout
	Data     map[string]interface{}
	Head     Head
	Error    string
	ErrorMsg string
}

func (BaseHandle) DefaultLayout() layout.Layout {
	return &layout.BaseLayout{}
}

func (this *BaseHandle) SetLayout(l layout.Layout) {
	this.layout = l
}

func (this BaseHandle) New() Handle {
	c := this
	c.Head = new(BaseHead).New()
	return &c
}

func (this *BaseHandle) Prepare(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println(`base Prepare`)
	return true
}

func (this *BaseHandle) SetUri(s string) {
	this.Uri = s
	// fmt.Println(`set uri =`, this)
}

func (this *BaseHandle) SetStartTime(t time.Time) {
	this.Time = t
}

func (this *BaseHandle) Get(r *http.Request) {
	// fmt.Println(r.URL.Query())
	// this.Data = make(map[string]interface{})
}

func (this *BaseHandle) Parse() *bytes.Buffer {
	if this.layout == nil {
		// fmt.Println(`new Parse`)
		this.layout = this.DefaultLayout()
	}

	if this.Data == nil {
		this.Data = make(map[string]interface{})
	}

	this.Data[`_head`] = this.Head.Export()
	this.Data[`_time`] = this.Time
	// fmt.Println(this.Data)

	if this.Error != `` {
		this.Uri = `/error/500`
		e := make(map[string]string)
		this.Data[`_error`] = e

		e[`title`] = this.Error
		e[`msg`] = this.ErrorMsg
	}

	this.layout.SetPath(this.Uri)
	this.layout.SetData(this.Data)
	return this.layout.Parse()
}
