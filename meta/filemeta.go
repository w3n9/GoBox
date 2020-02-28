package meta

import "time"

type FileMeta struct{
	FileName string`json:"filename"`
	FileHash string`json:"fileHash"`
	Size int`json:"size"`
	Location string`json:"location"`
	CreateAt time.Time`json:"createAt"`
}

var fileMetaSet map[string]FileMeta

func init(){
	fileMetaSet=make(map[string]FileMeta)
}

func UpdateFileMeta(fileMeta FileMeta){
	fileMetaSet[fileMeta.FileHash]=fileMeta
}
func DeleteFileMeta(fileHash string)FileMeta{
	res:=fileMetaSet[fileHash]
	delete(fileMetaSet,fileHash)
	return res
}
func GetFileMeta(fileHash string)(fileMeta FileMeta,ok bool){
	val,ok:=fileMetaSet[fileHash]
	return val,ok
}