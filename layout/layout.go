package layout

import (
	"bytes"
	"fmt"
	"html/template"
	"sync/atomic"
	"time"
)

var (
	bCache             = false
	connectCount int64 = 0
)

type Layout interface {
	SetFunc(f template.FuncMap)
	SetPath(string)
	SetData(interface{})
	Parse() *bytes.Buffer
}

type BaseLayout struct {
	Path     string
	Data     interface{}
	TplCache *template.Template
	Buffer   *bytes.Buffer
	FuncMap  template.FuncMap
}

func (this *BaseLayout) SetFunc(f template.FuncMap) {
	fmt.Println(`call SetFunc`)
	this.FuncMap = f
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

	this.ParseBaseTpl(`head`)
	this.ParseBaseTpl(`nav`)

	this.Buffer.WriteString("\n<div class=\"main container\">\n")
	this.ParseTpl(`tpl` + this.Path + `.html`)
	this.Buffer.WriteString("</div>\n")

	//time.Sleep(123 * time.Millisecond)

	t := this.Data.(map[string]interface{})[`_time`].(time.Time)
	td := float64(time.Now().Sub(t).Nanoseconds())/1000000 + 0.0005
	this.Data.(map[string]interface{})[`_stime`] = fmt.Sprintf(`%.03f`, td)
	this.Data.(map[string]interface{})[`_count`] = atomic.AddInt64(&connectCount, 1)
	this.ParseBaseTpl(`foot`)

	return this.Buffer
}

func (this *BaseLayout) ParseTpl(file string) {

	var tpl *template.Template
	var err error

	if bCache {
		tpl = TplPool.Get(file)
	} else {
		if this.FuncMap != nil {
			tpl, err = template.New(`k`).Option(`missingkey=zero`).Funcs(this.FuncMap).ParseFiles(file)
		} else {
			tpl, err = template.ParseFiles(file)
		}
		if err != nil {
			fmt.Println(`template error`, err)
			return
		}
	}
	if tpl == nil {
		return
	}

	if this.FuncMap != nil {
		err = tpl.ExecuteTemplate(this.Buffer, `main`, this.Data)
	} else {
		err = tpl.Execute(this.Buffer, this.Data)
	}
	if err != nil {
		fmt.Println(`file`, file, `tpl.Execute error:`, err)
	}
}

func (this *BaseLayout) ParseBaseTpl(name string) {
	BaseTpl.ExecuteTemplate(this.Buffer, name, this.Data)
}
