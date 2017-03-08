package layout

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"time"
)

var (
	bCache = true
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

	this.ParseBaseTpl(`head`)
	this.ParseBaseTpl(`nav`)
	this.ParseTpl(`tpl` + this.Path + `.html`)

	//time.Sleep(123 * time.Millisecond)

	t := this.Data.(map[string]interface{})[`_time`].(time.Time)
	td := float64(time.Now().Sub(t).Nanoseconds())/1000000 + 0.0005
	this.Data.(map[string]interface{})[`_stime`] = fmt.Sprintf(`%.03f`, td)
	this.ParseBaseTpl(`foot`)

	return this.Buffer
}

func formatCommas(num string) string {
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for {
		formatted := re.ReplaceAllString(num, "$1,$2")
		if formatted == num {
			return formatted
		}
		num = formatted
	}
}

func (this *BaseLayout) ParseTpl(file string) {

	var tpl *template.Template
	var err error

	if bCache {
		tpl = TplPool.Get(file)
	} else {
		tpl, err = template.ParseFiles(file)
		if err != nil {
			fmt.Println(`template error`, err)
			return
		}
	}
	if tpl == nil {
		return
	}

	tpl.Execute(this.Buffer, this.Data)
}

func (this *BaseLayout) ParseBaseTpl(name string) {
	BaseTpl.ExecuteTemplate(this.Buffer, name, this.Data)
}
