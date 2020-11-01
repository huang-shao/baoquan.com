package models

import (
	"bytes"
	"encoding/gob"
)

//数据上链的数据结构

type CertRecord struct {
	CertHash []byte
	CertId []byte
	CertAuthor string
	Phone string
	AuthorCard string
	FileName string
	FileSize int64
	CertTime int64
}

func (c CertRecord) SerializeRecord() ([]byte,error) {
	buff:=new(bytes.Buffer)
	err:=gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(),err

}
func DeSerializeRecord(data []byte)  (*CertRecord,error) {
	var certRecord *CertRecord
	err:=gob.NewDecoder(bytes.NewReader(data)).Decode(&certRecord)
	return certRecord,err
}
