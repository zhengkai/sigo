package layout

import (
	"bytes"
	// "fmt"
	"html/template"
)

type Layout interface {
	SetPath(string)
	SetData(interface{})
	Parse() *bytes.Buffer
}

type BaseLayout struct {
	Path     string
	Data     interface{}
	TplCache *template.Template
	Buffer   *bytes.Buffer
}

func (this *BaseLayout) SetData(d interface{}) {
	this.Data = d
}

func (this *BaseLayout) SetPath(s string) {
	this.Path = s
}

func (this *BaseLayout) Parse() *bytes.Buffer {

	if this.Buffer == nil {
		// fmt.Println(`buf init`)
		this.Buffer = new(bytes.Buffer)
	} else {
		this.Buffer.Reset()
	}

	this.ParseTpl(`tpl/head.html`)
	this.ParseTpl(`tpl/nav.html`)
	this.ParseTpl(`tpl` + this.Path + `.html`)
	this.ParseTpl(`tpl/foot.html`)

	return this.Buffer
}

func (this *BaseLayout) ParseTpl(s string) {

	// tpl, _ := template.ParseFiles(s)
	tpl := TplPool.Get(s)
	if tpl == nil {
		return
	}

	tpl.Execute(this.Buffer, this.Data)
}
