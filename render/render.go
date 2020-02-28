package render

import (
	"html/template"
	"net/http"
)
type Data map[string]interface{}
func HTML(w http.ResponseWriter,fileName string,data Data){
	tpl:=template.New("base")
	_, err := tpl.ParseFiles("./template/common/base.html", fileName)
	if err != nil {
		panic(err)
	}
	_ = tpl.Execute(w, data)
}
func HTMLWithBase(w http.ResponseWriter,data Data,base string,fileNames ...string){
	tpl, err := template.ParseFiles(fileNames...)
	if err != nil {
		panic(err)
	}
	_ = tpl.ExecuteTemplate(w, base, data)
}
