package layout

import (
	"bytes"
	"fmt"
	"html/template"
)

type Layout interface {
	SetPath(string)
	Parse() *bytes.Buffer
}

type BaseLayout struct {
	path string
}

func (this *BaseLayout) SetPath(s string) {
	fmt.Println(`this SetTpl =`, s)
	this.path = s
}

func (this *BaseLayout) Parse() *bytes.Buffer {

	buf := new(bytes.Buffer)

	tpl, _ := template.ParseFiles(`tpl/head.html`)
	tpl.Execute(buf, nil)

	tpl, _ = template.ParseFiles(`tpl/nav.html`)
	tpl.Execute(buf, nil)

	tpl, _ = template.ParseFiles(`tpl/foot.html`)
	tpl.Execute(buf, nil)

	return buf
}
