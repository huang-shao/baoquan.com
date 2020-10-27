package main

import (
	"baoquan_ruanda/blockchain"
	"baoquan_ruanda/db_baoquan"
	_ "baoquan_ruanda/routers"
	"fmt"
	"github.com/astaxie/beego"
)

var BUCKET_NAME  ="blocks"

func main() {





	block:=blockchain.CreateGenesisBlock()
	fmt.Println(block)
	fmt.Printf("区块hash值:%x\n",block.Hash)
	fmt.Printf("区块hash值:%d\n",block.Nonce)
	//db,err:=bolt.Open("chain.db",0600,nil)
	//if err!=nil {
	//	panic(err.Error())
	//}
	//defer db.Close()
	////操作chain.db数据库文件
	//db.Update(func(tx *bolt.Tx) error {
	//	var tong *bolt.Bucket
	//	tong=tx.Bucket([]byte(BUCKET_NAME))
	//	if tong==nil {
	//		tong,err=tx.CreateBucket([]byte(BUCKET_NAME))
	//		if err!=nil {
	//			return err
	//		}
	//	}
	//	//先查看获取桶中
	//	lastBlock:=tong.Get([]byte("lasthash"))
	//	blockHash,err:=block.Serialize()
	//	if err!=nil {
	//		return err
	//
	//	}
	//	if lastBlock==nil {
	//		tong.Put(block.Hash,blockHash)
	//		tong.Put([]byte("lasthash"),blockHash)
	//	}
	//	//tong.Put([]byte("0"),[]byte("a"))
	//	return nil
	//})




	return
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/login_js","./static/login_js")
	beego.SetStaticPath("/login_css","./static/login_css")
	beego.SetStaticPath("/login_img","./static/login_img")
	beego.SetStaticPath("/use_js","./static/use_js")
	beego.SetStaticPath("/use_css","./static/use_css")
	beego.SetStaticPath("/use_img","./static/use_img")
	db_baoquan.Init()
	beego.Run()
}

