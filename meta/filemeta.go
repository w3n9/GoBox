package meta

import (
	"GoBox/db"
	"fmt"
	"time"
)

type FileMeta struct{
	FileName string`json:"filename"`
	FileHash string`json:"fileHash"`
	Size int`json:"size"`
	Location string`json:"location"`
	CreateAt time.Time`json:"createAt"`
	UpdateAt time.Time`json:"updateAt"`
}


func UpdateFileMeta(fileMeta FileMeta)bool{
	stmt, err := db.Conn.Prepare(
		"insert into meta_file(file_sha1,file_name,file_size,file_addr,update_at,status)" +
			"value (?,?,?,?,?,'0')")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(fileMeta.FileHash, fileMeta.FileName, fileMeta.Size, fileMeta.Location,
		time.Now().In(time.UTC))
	if err != nil {
		panic(err)
	}
	rows,_:=res.RowsAffected()
	if rows<1{
		return false
	}else{
		return true
	}
}
//func DeleteFileMeta(fileHash string)FileMeta{
//	res:=fileMetaSet[fileHash]
//	delete(fileMetaSet,fileHash)
//	return res
//}
func GetFileMeta(fileHash string)(fileMeta FileMeta,ok bool){
	meta:=db.FileMeta{

	}
	stmt, _ := db.Conn.Prepare("select file_sha1,file_name,file_size,file_addr,create_at,update_at from meta_file where file_sha1=? and status=0 limit 1")
	defer stmt.Close()
	res:= stmt.QueryRow(fileHash)
	_ = res.Scan(&meta.FileSha1, &meta.FileName, &meta.FileSize, &meta.FileAddr,&(meta.CreateAt),&(meta.UpdateAt))
	fmt.Println(meta.CreateAt)
	fileMeta=FileMeta{
		FileName: meta.FileName,
		FileHash: meta.FileSha1,
		Size:     int(meta.FileSize.Int64),
		Location: meta.FileAddr,
		CreateAt: meta.CreateAt,
		UpdateAt: meta.UpdateAt,
	}
	return fileMeta,true
}