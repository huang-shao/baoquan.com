package controllers

import (
	"baoquan_ruanda/blockchain"
	"baoquan_ruanda/models"
	"baoquan_ruanda/util"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)

type FileUploadController struct {
	beego.Controller
}

func (f *FileUploadController) Get() {
	phone := f.GetString("name")
	f.Data["Phone"] = phone
	f.TplName = "storage.html"
}
func (this *FileUploadController) Post() {
	fileTitle := this.Ctx.Request.PostFormValue("upload_title")
	phone := this.Ctx.Request.PostFormValue("huangxunlin")
	//	fmt.Println("自定义的文件标题:",fileTitle)
	f, h, err := this.GetFile("myfile") //获取上传的文件
	if err != nil {
		this.Ctx.WriteString("获取失败。")
		return
	}
	defer f.Close()
	fmt.Println("自定义的文件标题:", fileTitle)
	fmt.Println("文件名称：", h.Filename)
	fmt.Println("文件的大小：", h.Size)
	fmt.Println(f)
	//this.Ctx.WriteString("解析到上传文件："+h.Filename)

	//将文件保存在本地的一个目录中
	//文件全路径=路径+文件名+"."+扩展名
	uploadDir := "./static/img/" + h.Filename
	saveFile, err := os.OpenFile(uploadDir, os.O_RDWR|os.O_CREATE, 777)
	//创建一个writer:用于向硬盘上写一个文件
	writer := bufio.NewWriter(saveFile)

	file_size, err := io.Copy(writer, f)
	if err != nil {
		fmt.Println(err.Error())
		this.Ctx.WriteString("抱歉，保存的电子数据失败。")
		return
	}
	fmt.Println("拷贝的文件的大小是：", file_size)
	defer saveFile.Close()

	//计算文件的hash
	hashFile, err := os.Open(uploadDir)
	defer hashFile.Close()
	hash, err := util.MD5HashReader(hashFile)
	//md5Hash := md5.New()
	//fileBytes, _ := ioutil.ReadAll(f)
	//md5Hash.Write(fileBytes)
	//hashBytes := md5Hash.Sum(nil)
	//hash := hex.EncodeToString(hashBytes)

	//将上传 的记录保存到数据库里
	record := models.UploadRecord{}
	record.FileName = h.Filename
	record.FileSize = h.Size
	record.FileTitle = fileTitle
	record.CertTime = time.Now().Unix() //毫秒数
	record.FileCert = hash
	record.User_Name = phone //手机号

	_, err = record.SaveRecord()
	if err != nil {
		fmt.Println(err)
		this.Ctx.WriteString("抱歉，数据认证失败！")
		return
	}
	//将需要认证的文件hash值及个人实名信息保存到区块链上
	_, err = blockchain.CHAIN.SaveData([]byte(hash))
	if err != nil {
		this.Ctx.WriteString("抱歉，认证数据上链失败,请重试！")
		return
	}
	//从数据库中读取用户对应的所有认证数据记录
	records, err := models.QueryRecord(phone)
	if err != nil {
		this.Ctx.WriteString("抱歉，获取数据认证数据失败！")
		return
	}
	this.Data["Records"] = records
	this.Data["Name"] = phone
	this.TplName = "list_record.html"

	//this.Ctx.WriteString("成功获取到上传的文件！")

	//ext := path.Ext(h.Filename)
	//
	////验证后缀名是否符合要求
	//var AllowExtMap map[string]bool = map[string]bool{
	//	".jpg":true,
	//	".jpeg":true,
	//	".png":true,
	//}
	//if _,ok:=AllowExtMap[ext];!ok{
	//	this.Ctx.WriteString( "后缀名不符合上传要求" )
	//	return
	//}
	//创建目录
	//uploadDir := "static/img/" //+ time.Now().Format("2006/01/02/")
	//err = os.MkdirAll( uploadDir , 777)
	//if err != nil {
	//	this.Ctx.WriteString( fmt.Sprintf("%v",err) )
	//	return
	//}
	////构造文件名称
	//rand.Seed(time.Now().UnixNano())
	//randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	//hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )
	//
	//FileName := fmt.Sprintf("%x",hashName) + ext
	////this.Ctx.WriteString(  fileName )
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//fpath := uploadDir + FileName
	//defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	//err = this.SaveToFile("myfile", fpath)
	//if err != nil {
	//	this.Ctx.WriteString( fmt.Sprintf("%v",err) )
	//}
	//this.Ctx.WriteString( "上传成功~！！！！！！！" )
}
