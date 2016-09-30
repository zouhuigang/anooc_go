package main

import (
	"fmt"
	"net/http"
	"strconv"
	"io/ioutil"
	"os"
	"path"
	"log"
	"sync"
)

var ch chan int = make(chan int)
var qqinit int=26807293
func main() {
    var wg sync.WaitGroup //创建一个sync.WaitGroup
	TCount :=50 //并发
	//产生任务
   go func() {
    for i := 0; i < 200; i++ {
	     ch <- i
    }
    close(ch)
  }()

  //开始执行
   wg.Add(TCount)
   for i := 0; i < TCount; i++ {
     i := i
	 go func() {
      defer func() { wg.Done() }()

      fmt.Printf("工作者 %v 启动...\r\n", i)

      for task := range ch {
        func() {

          defer func() {
            err := recover()
            if err != nil {
              fmt.Printf("任务失败：工作者i=%v, task=%v, err=%v\r\n", i, task, err)
            }
          }()
		  
            geticon(task)
	        fmt.Printf("任务结果=%v ，工作者id=%v, task=%v\r\n",task*task,i,task)
        }()
      }

      fmt.Printf("工作者 %v 结束。\r\n", i)
    }()
	
   }
   
  //等待所有任务完成
  wg.Wait()
  print("全部任务结束")
}

//抓取头像17923391;19923392,20023392
func geticon(groupNumber int){
	number := groupNumber*1000+qqinit;
	end := (groupNumber+1)*1000+qqinit;
	for i:=number; i<end;i++ {
		 // 通过Itoa方法转换
		 str1 := strconv.Itoa(i)
		 urls :="http://qlogo3.store.qq.com/qzone/"+str1+"/"+str1+"/100"
		 downOneImg(urls,str1,groupNumber)
	}
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
	_, err = f.Write(body)
	
    defer f.Close()

}
