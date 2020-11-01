package main

import (
	"baoquan_ruanda/blockchain"
	"baoquan_ruanda/db_baoquan"
	_ "baoquan_ruanda/routers"
	"github.com/astaxie/beego"
)

//var BUCKET_NAME  ="blocks"

func main() {
	//首先，准备一条区块链
	blockchain.NewBlockChain()






	//1、实例化一个区块链实例
	//bc:=blockchain.NewBlockChain()
	//fmt.Printf("最新区块的hash值:%x\n",bc.LastHash)
	//blocks:=bc.QueryAllBlocks()
	//if len(blocks)==0 {
	//	fmt.Println("暂未查询到区块链数据")
	//	return
	//}
	//for _,block:=range blocks{
	//	fmt.Printf("高度:%d,哈希:%x,prev哈希%x\n",block.Height,block.Hash,block.PrevHash)
	//}



	//数据被保存在height为1的区块当中
	//block,err:=bc.SaveData([]byte("这里存储上链的数据信息"))
	//if err!=nil {
	//	fmt.Println(err.Error())
	//	return
	//
	//}
	//fmt.Printf("区块的高度:%d\n",block.Height)
	//fmt.Printf("区块的prevhash:%x\n",block.PrevHash)
	////fmt.Printf("新区块的hash:%d\n",block.Hash)
	//block1:=bc.QueryBlockByHeight(1)
	//if block1==nil {
	//	fmt.Println("输入错误！")
	//	return
	//}
	//	fmt.Println("区块的高度是:",block1.Height)
	//	fmt.Println("区块存储的信息是:",string(block1.Data))
	//
	//return

	//block:=blockchain.CreateGenesisBlock()
	//fmt.Println(block)
	//fmt.Printf("区块hash值:%x\n",block.Hash)
	//fmt.Printf("区块hash值:%d\n",block.Nonce)
	//db,err:=bolt.Open("chain1.db",0600,nil)
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
	//}
	db_baoquan.Init()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/login_js","./static/login_js")
	beego.SetStaticPath("/login_css","./static/login_css")
	beego.SetStaticPath("/login_img","./static/login_img")
	beego.SetStaticPath("/use_js","./static/use_js")
	beego.SetStaticPath("/use_css","./static/use_css")
	beego.SetStaticPath("/use_img","./static/use_img")

	beego.Run()
}

