package handle

import (
	// "fmt"
	"html/template"
)

type BaseHead struct {
	Data   map[string]interface{}
	Unique map[string]bool
}

func (this BaseHead) New() Head {
	c := this
	c.Data = make(map[string]interface{})
	c.Unique = make(map[string]bool)
	return &c
}

func (this *BaseHead) AddJS(s string) {
	this.AddData(`js`, s)
}

func (this *BaseHead) AddCSS(s string) {
	this.AddData(`css`, s)
}

func (this *BaseHead) AddMeta(s string) {
	this.AddData(`meta`, s)
}

func (this *BaseHead) AddData(k, v string) {
	if _, ok := this.Unique[v]; ok {
		return
	}

	if k == `meta` {
		if _, ok := this.Data[k]; !ok {
			this.Data[k] = []template.HTML{}
		}
		this.Data[k] = append(this.Data[k].([]template.HTML), template.HTML(v))
		return
	}

	if _, ok := this.Data[k]; !ok {
		this.Data[k] = []string{}
	}
	this.Data[k] = append(this.Data[k].([]string), v)
}

func (this *BaseHead) Export() map[string]interface{} {
	// fmt.Println(this.Data)
	return this.Data
}
