package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"server/gather/lib"
	"strconv"
	"sync"
)

var qqinit int = 26807293
var ch chan int = make(chan int)

func main() {
	var wg sync.WaitGroup //创建一个sync.WaitGroup
	TCount := 20          //并发
	//产生任务
	go func() {
		for i := 0; i < 10; i++ {
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
					fmt.Printf("任务结果=%v ，工作者id=%v, task=%v\r\n", task*task, i, task)
				}()
			}

			fmt.Printf("工作者 %v 结束。\r\n", i)
		}()

	}

	//等待所有任务完成
	wg.Wait()
	print("全部任务结束")

}

func geticon(groupNumber int) {
	number := groupNumber*1000 + qqinit
	end := (groupNumber+1)*1000 + qqinit
	for i := number; i < end; i++ {
		// 通过Itoa方法转换
		str1 := strconv.Itoa(i)
		urls := "http://cgi.find.qq.com/qqfind/buddy/search_v3?num=1&page=0&sessionid=0&keyword=" + str1 + "&agerg=0&sex=0&firston=1&video=0&country=1&province=31&city=0&district=0&hcountry=1&hprovince=0&hcity=0&hdistrict=0&online=1&ldw=1506481715"
		downOneImg(urls, str1, groupNumber)
	}
}

//开始爬取
func downOneImg(imgUrl string, id string, groupNumber int) {

	//防止程序出错
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()

	//模拟请求
	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", imgUrl, nil)

	reqest.Header.Set("Host", "cgi.find.qq.com")
	reqest.Header.Set("Content-Length", "181")
	reqest.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	reqest.Header.Set("Origin", "http://find.qq.com")
	reqest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.59 QQ/8.5.18600.201 Safari/537.36")
	reqest.Header.Set("Referer", "http://find.qq.com/index.html?version=1&im_version=5491&width=910&height=610&search_target=0")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate")
	reqest.Header.Set("Accept-Language", "en-us,en")
	reqest.Header.Set("Cookie", "uin=o381902897;skey=ZKFI9LreOO;")

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

	//开始解析
	lib.JsonData(string(body), "1.csv")

}
