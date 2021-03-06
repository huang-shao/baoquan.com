package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Height int64
	TimeStamp int64
	Hash []byte
	Data []byte
	PrevHash []byte
	Version string
	Nonce int64

}

//生成创世区块，返回区块信息
func CreateGenesisBlock() Block  {
	block:=NewBlock(0,[]byte{},[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return block
}

func NewBlock(height int64,data []byte,preHash []byte)(Block)  {
	block:=Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		//Hash:      nil,
		Data:      data,
		PrevHash:  preHash,
		Version:   "0x01",
		//Nonce:     0,
	}
	pow :=NewPoW(block)
	blockHash,nonce:=pow.Run()
	block.Nonce=nonce
	block.Hash=blockHash

	//heightBytes,_:=util.IntToBytes(block.Height)
	//timeBytes,_:=util.IntToBytes(block.TimeStamp)
	//versionBytes:=util.StringToBytes(block.Version)
	//nonceBytes,_:=util.IntToBytes(block.Nonce)
	//blockBytes:=bytes.Join([][]byte{
	//	heightBytes,
	//	timeBytes,
	//	data,
	//	preHash,
	//	versionBytes,
	//	nonceBytes,
	//},[]byte{})
	//
	//block.Hash=util.SHA256Hash(blockBytes)
	return block
}
//区块的序列化
func (bk Block)Serialize() ([]byte,error) {
	buff:=new(bytes.Buffer)
	err:=gob.NewEncoder(buff).Encode(bk)
	if err!=nil {
		return nil,err

	}
	return buff.Bytes(),nil
}




//区块的反序列化
func DeSerialize(data []byte)(*Block,error)  {
	var block  Block
	err:=gob.NewDecoder(bytes.NewBuffer(data)).Decode(&block)
	if err!=nil {
		return nil,err
	}
	return &block,nil

}

