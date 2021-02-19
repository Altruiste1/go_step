package utools

import "os"

func CreateFile(fileName string,filePath string)error{
	_,err:=os.Create(filePath+fileName)
	if err!=nil{
		return err
	}
	return nil
}

func OpenFile(fileName string,fileMode os.FileMode)  {
	os.OpenFile(fileName,0,fileMode)

}