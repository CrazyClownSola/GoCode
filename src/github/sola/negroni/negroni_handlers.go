package negroni

import (
	"log"
	"net/http"
)

//中间件
type MiddleWare struct {
	IsBool bool
	Name   string
}

func NewMiddleWare() *MiddleWare {
	return &MiddleWare{true, "STR"}
}

func (m *MiddleWare) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandleFunc) {

}
