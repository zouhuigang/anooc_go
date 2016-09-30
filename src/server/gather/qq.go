package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"path"
	"encoding/csv"
	"encoding/json"
)

func main(){

    Spy("http://cgi.find.qq.com/qqfind/buddy/search_v3?num=20&page=0&sessionid=0&keyword=285179335&agerg=0&sex=0&firston=1&video=0&country=1&province=31&city=0&district=0&hcountry=1&hprovince=0&hcity=0&hdistrict=0&online=1&ldw=2109567036")
	fmt.Println("抓取qq")
}


//开始爬取
func Spy(url string) {

    //防止程序出错
    defer func() {
        if r := recover(); r != nil {
            log.Println("[E]", r)
        }
    }()

    //模拟请求
    client := &http.Client{}
    reqest, _ := http.NewRequest("POST", url, nil)

	reqest.Header.Set("Host","cgi.find.qq.com")
	reqest.Header.Set("Content-Length","181")
	reqest.Header.Set("Accept","application/json, text/javascript, */*; q=0.01")
	reqest.Header.Set("Content-Type","application/x-www-form-urlencoded; charset=UTF-8")
	reqest.Header.Set("Origin","http://find.qq.com")
	reqest.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.59 QQ/8.5.18600.201 Safari/537.36")
	reqest.Header.Set("Referer","http://find.qq.com/index.html?version=1&im_version=5491&width=910&height=610&search_target=0")
	reqest.Header.Set("Accept-Encoding","gzip,deflate")
	reqest.Header.Set("Accept-Language","en-us,en")
	reqest.Header.Set("Cookie","itkn=1699827159; pt2gguin=o0381902897; uin=o0381902897; skey=@tAX8jPGxY; ptisp=ctc; RK=FRfzh5jHPE; ptcz=ea6c60baa547c86e89b81ae3de55ae61265067e1952c7ad7a74571976021ed69; pgv_info=ssid=s2400012874; pgv_pvid=2902014596; o_cookie=381902897")


	 
    response, err := client.Do(reqest)
   

    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        return
    }
	
	 body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
		return
    }
	defer response.Body.Close()
	
	 //写入文件
	fileName :="12.txt"
	os.MkdirAll(path.Dir(fileName), os.ModePerm)
    f, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
	//写入图信息
	_, err = f.Write(body)
	
    defer f.Close()
	
	

    
}


//保存信息为cvs格式的，方便导入数据库
func savecvs(){
	f, err := os.Create("test.csv")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
    w := csv.NewWriter(f)
    w.Write([]string{"nick","email","mobile"})
    w.Write([]string{"1","张三","23"})
    w.Write([]string{"2","李四","24"})
    w.Write([]string{"3","王五","25"})
    w.Write([]string{"4","赵六","26"})
    w.Flush()
}