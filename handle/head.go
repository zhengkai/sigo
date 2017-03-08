package handle

import (
// "fmt"
)

type Head struct {
	Data   map[string]interface{}
	IsInit bool
}

func (this *Head) Init() {
	if this.IsInit {
		// fmt.Println(`head init skip`)
		return
	}
	d := make(map[string]interface{})
	d[`js`] = make(map[string]bool)
	d[`css`] = make(map[string]bool)
	d[`meta`] = make(map[string]bool)
	this.Data = d
	this.IsInit = true
	// fmt.Println(`head init`)
}

func (this *Head) AddJS(s string) {
	this.Init()
	this.Data[`js`].(map[string]bool)[s] = true
}

func (this *Head) AddCSS(s string) {
	this.Init()
	this.Data[`css`].(map[string]bool)[s] = true
}
func (this *Head) AddMeta(s string) {
	this.Init()
	this.Data[`meta`].(map[string]bool)[s] = true
}

func (this *Head) Export() map[string]interface{} {
	this.Init()
	return this.Data
}
