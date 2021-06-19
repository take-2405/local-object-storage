package decode

import (
	"encoding/base64"
	"os"
)

////エンコード
//func Encode() string {
//
//	file, _ := os.Open("./../images/IMG_8173.png")
//	defer file.Close()
//
//	fi, _ := file.Stat() //FileInfo interface
//	size := fi.Size()    //ファイルサイズ
//
//	data := make([]byte, size)
//	file.Read(data)
//
//	return base64.StdEncoding.EncodeToString(data)
//}
//
////デコード
//func Decode(str string) {
//	data, _ := base64.StdEncoding.DecodeString(str) //[]byte
//
//	file, _ := os.Create("./../images/encode_and_decord.png")
//	defer file.Close()
//	file.Write(data)
//}
//
func Decode(info,fileName string)error {
	var err error
	data, err := base64.StdEncoding.DecodeString(info) //[]byte
	if err!=nil{
		return err
	}
	file, err := os.Create(fileName)
	if err!=nil{
		return err
	}
	defer file.Close()

	_,err = file.Write(data)
	if err!=nil{
		return err
	}
	return err
}
