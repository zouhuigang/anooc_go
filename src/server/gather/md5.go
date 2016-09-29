/*
删除抓图片时抓下来的重复的图
md5 9e981dc7788599d9372c8381d776c554
md5 03b42d62450ffcb05428b958512bea22
*/

package main                                                                       

import (                                                                           
        "crypto/md5"  
        "encoding/hex"		
        "fmt"                                                                       
        "io"                                                                                                                                            
        "os"  
		"path/filepath"
)                                                                                   

//md5数组
var imgmd5 = "03b42d62450ffcb05428b958512bea22"
var imgarr = [...]string {
		"03b42d62450ffcb05428b958512bea22",
		"9e981dc7788599d9372c8381d776c554"}

var ch chan int = make(chan int)

func main() {                                                                       
        //testFile :="D:/www/test-go/src/server/gather/pic23409330/999/24408942.jpg"                                                     
        //fileMd5 := getFileMd5(testFile)
		//fmt.Printf(fileMd5)
		
		//fmt.Printf("ddd")
		//os.Exit(1)--退出
		
		testFile :="D:/www/test-go/src/server/gather/pic9962866"
		for i :=0; i <50; i++ {
			go getFilelist(testFile) 
		}
		
		for i :=0; i <50; i++ {
			<-ch
		}
		
		fmt.Printf("success")
       		
}


//得到文件图片的md5值
func getFileMd5(pathFile string) string{
       
		  file, _ := os.Open(pathFile)		  
          md5h := md5.New()                                                   
          io.Copy(md5h, file) 
		  cipherStr := md5h.Sum(nil)	
		  defer file.Close()
		  //加密输出md5值
          return hex.EncodeToString(cipherStr)

}

//搜索文件夹下的所有文件目录
func getFilelist(path string) {
        err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
                if ( f == nil ) {return err}
                if f.IsDir() {return nil}
                //println(path)
				delFile(path)
                return nil
        })
        if err != nil {
                fmt.Printf("filepath.Walk() returned %v\n", err)
        }
		ch <-1
}

//删除文件
func delFile(path string){
	fileMd5 := getFileMd5(path)
	 if  matchingFileMd5(fileMd5) {
		fmt.Printf("已找到待删除文件"+path+"\n")
		err := os.Remove(path)
        if err != nil {
           //输出错误详细信息
           fmt.Printf("%s", err)
        } 
	
	} else {
	   //fmt.Printf("未找到待删除文件\n")
	}
	
	

}

//判断找到的文件是否和定义的文件md5一样
func matchingFileMd5(filemd5 string) bool {
		for _,v := range imgarr{
			 if  filemd5 == v {
			      return true
			 }
		}	
		return false
}
