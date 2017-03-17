package handle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/zhengkai/sigo/layout"
)

type ContentType int

const (
	Html ContentType = iota
	Json
	Jsonp
	Text
	Markdown
)

type BaseHandle struct {
	Time        time.Time
	Uri         string
	layout      layout.Layout
	Data        interface{}
	Head        Head
	Error       string
	ErrorMsg    string
	ContentType ContentType
	Cookie      []*http.Cookie
	W           http.ResponseWriter
	R           *http.Request
}

func (BaseHandle) DefaultLayout() layout.Layout {
	return &layout.BaseLayout{}
}

func (this *BaseHandle) AddCookie(*http.Cookie) {
}

func (this *BaseHandle) SetLayout(l layout.Layout) {
	this.layout = l
}

func (this BaseHandle) New() Handle {
	c := this
	c.Head = new(BaseHead).New()
	return &c
}

func (this *BaseHandle) Prepare() bool {
	fmt.Println(`base Prepare`)
	return true
}

func (this *BaseHandle) SetUri(s string) {
	this.Uri = s
}

func (this *BaseHandle) SetHttp(w http.ResponseWriter, r *http.Request) {
	this.W = w
	this.R = r
}

func (this *BaseHandle) SetStartTime(t time.Time) {
	this.Time = t
}

func (this *BaseHandle) Parse() {
}

func (this *BaseHandle) Output() *bytes.Buffer {

	switch this.ContentType {
	case Json:
		return this.OutputJSON()
	case Html:
		return this.OutputHTML()
	default:
		buf := new(bytes.Buffer)
		buf.WriteString(`no support type ` + strconv.Itoa(int(this.ContentType)))
		return buf
	}
}

func (this *BaseHandle) OutputHTML() *bytes.Buffer {

	if this.layout == nil {
		// fmt.Println(`new Parse`)
		this.layout = this.DefaultLayout()
	}
	if this.Data == nil {
		this.Data = make(map[string]interface{})
	}
	this.Data.(map[string]interface{})[`_head`] = this.Head.Export()
	this.Data.(map[string]interface{})[`_time`] = this.Time
	// fmt.Println(this.Data)

	if this.Error != `` {
		this.Uri = `/error/500`
		e := make(map[string]string)
		this.Data.(map[string]interface{})[`_error`] = e

		e[`title`] = this.Error
		e[`msg`] = this.ErrorMsg
	}

	this.layout.SetPath(this.Uri)
	this.layout.SetData(this.Data)
	return this.layout.Parse()
}

func (this *BaseHandle) Redirect(uri string) {
	http.Redirect(this.W, this.R, uri, 302)
}

func (this *BaseHandle) OutputJSON() *bytes.Buffer {
	this.W.Header().Set(`Content-Type`, `application/json; charset=utf-8`)
	buf := new(bytes.Buffer)
	b, _ := json.Marshal(this.Data)
	// fmt.Println(`json`, this.Data)
	buf.Write(b)
	return buf
}
