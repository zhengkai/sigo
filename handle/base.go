package handle

import (
	"bytes"
	"time"
	// "fmt"
	"net/http"

	"github.com/zhengkai/sigo/layout"
)

type BaseHandle struct {
	time   time.Time
	uri    string
	layout layout.Layout
	Data   map[string]interface{}
	Head   Head
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

func (this *BaseHandle) SetUri(s string) {
	this.uri = s
	// fmt.Println(`set uri =`, this)
}

func (this *BaseHandle) SetStartTime(t time.Time) {
	this.time = t
}

func (this *BaseHandle) Get(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL.Query())
	// this.Data = make(map[string]interface{})
}

func (this *BaseHandle) Parse() *bytes.Buffer {
	if this.layout == nil {
		// fmt.Println(`new Parse`)
		this.layout = this.DefaultLayout()
	}
	this.Data[`_head`] = this.Head.Export()
	this.Data[`_time`] = this.time
	// fmt.Println(this.Data)

	this.layout.SetPath(this.uri)
	this.layout.SetData(this.Data)
	return this.layout.Parse()
}
