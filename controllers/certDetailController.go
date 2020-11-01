package controllers

import (
	"baoquan_ruanda/blockchain"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

//证书详细信息查看页面控制器
type CertDetailController struct {
	beego.Controller
}

func (c *CertDetailController) Get()  {
	//0、获取前端页面get请求时携带的cert_id数据
		certId:=c.GetString("cert_id")
		fmt.Println("需要认证的id：",certId)
	//①、准备数据：根据cert_id到区块链上查询具体信息，获取区块信息
		block,err:=blockchain.CHAIN.QueryBlockByCertId([]byte(certId))
		if err!=nil{
			c.Ctx.WriteString("链上数据查询失败！")
			return
		}
		//查询未遇到错误，分为两种情况：查到与未查到
	if block==nil {
		c.Ctx.WriteString("抱歉，未查询到数据，请重试！")
		return
	}
	//certId=hex.EncodeToString(block.Data)
	//大写转换
	c.Data["CertId"]=strings.ToUpper(string(block.Data))
	//②、跳转页面
	c.TplName="cert_detail.html"

}
