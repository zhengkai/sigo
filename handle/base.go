package handle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
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
	StopStatus  int
	TplFuncMap  template.FuncMap
}

func (BaseHandle) DefaultLayout() layout.Layout {
	return &layout.BaseLayout{}
}

func (this *BaseHandle) AddCookie(*http.Cookie) {
}

func (this *BaseHandle) SetTplFunc(name string, fn interface{}) {
	if this.TplFuncMap == nil {
		this.TplFuncMap = make(template.FuncMap)
	}
	this.TplFuncMap[name] = fn
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

func (this *BaseHandle) StopByStatus(status int) {
	this.W.WriteHeader(status)
	this.StopStatus = status
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

func (this *BaseHandle) Output() {

	if this.StopStatus > 0 {
		this.W.Header().Set(`Content-Type`, `text/plain; charset=utf-8`)
		this.W.Write([]byte(`HTTP ` + strconv.Itoa(this.StopStatus) + ` ` + http.StatusText(this.StopStatus)))
		return
	}

	var buf *bytes.Buffer

	switch this.ContentType {
	case Json:
		buf = this.OutputJSON()
	case Html:
		buf = this.OutputHTML()
	default:
		buf = new(bytes.Buffer)
		buf.WriteString(`no support type ` + strconv.Itoa(int(this.ContentType)))
	}

	this.W.Write(buf.Bytes())
}

func (this *BaseHandle) CheckPost() bool {
	// fmt.Println(`referer`, this.R.Header.Get(`Referer`))
	if this.R.Method != http.MethodPost {
		this.StopByStatus(http.StatusMethodNotAllowed)
		return false
	}
	return true
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
		this.W.WriteHeader(http.StatusInternalServerError)
		this.Uri = `/error/500`
		e := make(map[string]string)
		this.Data.(map[string]interface{})[`_error`] = e

		e[`title`] = this.Error
		e[`msg`] = this.ErrorMsg
	}

	this.layout.SetPath(this.Uri)
	this.layout.SetData(this.Data)

	if this.TplFuncMap != nil {
		this.layout.SetFunc(this.TplFuncMap)
	}

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
