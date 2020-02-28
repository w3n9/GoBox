package handler

import (
	"GoBox/config"
	"GoBox/meta"
	"GoBox/render"
	"GoBox/util"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)
func GetFileMetaHandler(w http.ResponseWriter,r *http.Request){

}
func UploadFileHandler(w http.ResponseWriter,r *http.Request){
	if r.Method==http.MethodGet{
		params:=r.URL.Query()
		if val,ok:=params["hash"];ok&&val[0]!=""{
			if fileMeta,ok:=meta.GetFileMeta(val[0]);ok{
				jsonStr, err := json.Marshal(fileMeta)
				_, err = io.WriteString(w, string(jsonStr))
				if err != nil {
					panic(err)
				}
			}else{
				w.WriteHeader(http.StatusNotFound)
			}
		} else{
			render.HTML(w,"./template/upload/index.html",render.Data{
				"title":"file/upload",
			})
		}
	}else if r.Method==http.MethodPost{
		file, header, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()
		fileMeta:=meta.FileMeta{}
		fileMeta.FileName=header.Filename
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		fileMeta.FileHash=util.Sha1(bytes)
		fileMeta.Location=config.Get().Storage.Path+fileMeta.FileHash
		newFile,err:=os.Create(fileMeta.Location)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer newFile.Close()
		size, err := io.Copy(newFile, file)
		fileMeta.Size=int(size)
		fileMeta.CreateAt=time.Now()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		meta.UpdateFileMeta(fileMeta)
		http.Redirect(w,r,"/file/success",http.StatusFound)
	}else if r.Method==http.MethodPut{
		all, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fileMeta:=meta.FileMeta{}
		err = json.Unmarshal(all, &fileMeta)
		if err != nil {
			panic(err)
		}
		if oldFileMeta,ok:=meta.GetFileMeta(fileMeta.FileHash);ok{
			oldFileMeta.FileName=fileMeta.FileName
			meta.UpdateFileMeta(oldFileMeta)
			_, _ = io.WriteString(w, "update ok")
		}else{
			_, _ = io.WriteString(w, "hash error")
		}

	}else if r.Method==http.MethodDelete{
		hash:=r.URL.Query().Get("hash")
		fileMeta := meta.DeleteFileMeta(hash)
		err := os.Remove(fileMeta.Location)
		if err != nil {
			panic(err)
		}
		_, _ = io.WriteString(w, "del ok")
	} else{
		w.WriteHeader(http.StatusBadRequest)
	}
}
func UploadFileSuccessHandler(w http.ResponseWriter,r *http.Request){
	render.HTMLWithBase(w,render.Data{
		"title":"",
	},"base","./template/common/base.html","./template/upload/success.html")
}
