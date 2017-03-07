package layout

import (
	// "fmt"
	"html/template"
	"sync"
)

var (
	TplPool = Tpl{Pool: make(map[string]*template.Template)}
)

type Tpl struct {
	Pool map[string]*template.Template
	Lock sync.Mutex
}

func (this *Tpl) Get(s string) *template.Template {

	if tpl, ok := this.Pool[s]; ok {
		return tpl
	}

	// fmt.Println(`new tpl`, s)

	this.Lock.Lock()
	defer this.Lock.Unlock()

	tpl, _ := template.ParseFiles(s)
	this.Pool[s] = tpl
	return tpl
}
