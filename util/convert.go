package util

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(num int64)([]byte,error)  {
	buff:=new(bytes.Buffer)
	err:=binary.Write(buff,binary.BigEndian,num)
	if err!=nil {
		return nil,err
		
	}
	return buff.Bytes(),nil


}

func StringToBytes(st string)[]byte  {
	return []byte(st)

}