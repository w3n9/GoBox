package main

import (
	"GoBox/config"
	"GoBox/handler"
	"net/http"
)

func main(){
	conf:=config.Get()
	mux:=http.NewServeMux()
	fileServer:=http.FileServer(http.Dir(conf.Web.StaticDirPrefix))
	mux.Handle(conf.Web.StaticUrlPrefix,http.StripPrefix(conf.Web.StaticUrlPrefix,fileServer))
	mux.HandleFunc("/file",handler.UploadFileHandler)
	mux.HandleFunc("/file/success",handler.UploadFileSuccessHandler)
	server:=http.Server{
		Addr:              ":8000",
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	_ = server.ListenAndServe()
}