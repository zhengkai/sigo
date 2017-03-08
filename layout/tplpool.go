package layout

import (
	"fmt"
	"html/template"
	"strings"
	"sync"
)

var (
	TplPool = Tpl{Pool: make(map[string]*template.Template)}
	funcMap = template.FuncMap{
		"HasSuffix": strings.HasSuffix,
	}
	BaseTpl, _ = template.New(`Base`).Funcs(funcMap).ParseFiles(`tpl/common.html`)
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

	tpl, err := template.ParseFiles(s)
	this.Pool[s] = tpl

	if err != nil {
		fmt.Println(`template error`, err)
	}

	return tpl
}
