package mux

import "net/http"

type customeHandler func(w http.ResponseWriter, r *http.Request)
type rules struct{ rulesMap map[string]customeHandler }

func (this *rules) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn, exit := this.rulesMap[r.URL.Path]
	if exit == false {
		http.NotFound(w, r)
	} else {
		fn(w, r)
	}
	//fmt.Fprintf(w, "hola")
}
func (this *rules) AddFun(rule string, fn customeHandler) {
	this.rulesMap[rule] = fn
}
func (this *rules) AddHandle(rule string, handle http.Handler) {
	this.rulesMap[rule] = handle.ServeHTTP
}
func CreateMux() *rules {
	newMap := make(map[string]customeHandler)
	return &rules{newMap}

}
