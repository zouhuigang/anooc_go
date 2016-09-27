package main

import (
	"fmt"
	"net/http"
	"strconv"
	"io/ioutil"
	"os"
	"path"
	"log"
)

var ch chan int = make(chan int)
var qqinit int=21417933
func main() {
	groupNumber :=200
	for i :=0;i<groupNumber;i++ {
		go geticon(i)
	}
	
	for i :=0;i<groupNumber;i++ {
		<-ch
	}
	
	fmt.Println("success")
}

//抓取头像17923391;19923392,20023392
func geticon(groupNumber int){
	number := groupNumber*10000+qqinit;
	end := (groupNumber+1)*10000+qqinit;
	for i:=number; i<end;i++ {
		 // 通过Itoa方法转换
		 str1 := strconv.Itoa(i)
		 urls :="http://qlogo3.store.qq.com/qzone/"+str1+"/"+str1+"/100"
		 downOneImg(urls,str1,groupNumber)
	}
	ch <- 1
}


//下载远程图片，无jpg,png等后缀的网址
func downOneImg(imgUrl string,id string,groupNumber int) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()
	
    req, err := http.NewRequest("GET", imgUrl, nil)
    if err != nil {
        panic(err)
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
		return
    }
	defer resp.Body.Close()
    //写入文件
	fileName :="pic"+strconv.Itoa(qqinit)+"/"+strconv.Itoa(groupNumber)+"/"+id+".jpg"
	os.MkdirAll(path.Dir(fileName), os.ModePerm)
    f, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
	//写入图信息
	_, err = f.Write([]byte(body))
	
    defer f.Close()

}
