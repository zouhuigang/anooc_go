
## linux/window运行golang程序
### 下载镜像

    docker pull daocloud.io/library/golang:1.7rc1


### linux:

编译不运行：
将当前目录作为一个 volume 挂载进容器，并把这个 volume 设置成工作目录，然后运行命令go build来编译应用：

    docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp daocloud.io/golang:1.7rc1 go build -v
会生成一个myapp的文件

运行软件

./myapp


----------

### windows:

编译windows版本：

    docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=windows -e GOARCH=386 daocloud.io/golang:1.7rc1 go build -v

会生成一个myapp.exe的文件

运行软件

点击运行


----------


编译多版本:

    $ docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.7rc1 bash
    $ for GOOS in darwin linux; do
    >   for GOARCH in 386 amd64; do
    >     go build -v -o myapp-$GOOS-$GOARCH
    >   done
    > done


#### http://139.196.16.67:8001/hello

# windows下的docker操作

### 配置挂载

将d:/mnt挂载进linux的虚拟机Oracle VM VirtualBox/vm中 /mnt-linux

vm虚拟机中的挂载命令：

    mount -t vboxsf mnt /mnt-linux

vm虚拟机挂载进docker容器中：

    docker run -i -t -v /mnt-linux:/usr/src/myapp golang:1.7rc1 /bin/bash

这样就把windwos中d盘下的d:/mnt目录挂载进了docker中的/usr/src/myapp


### 运用

测试main.go:

    package main

    import (
	    "net/http"
    )

    func SayHello(w http.ResponseWriter, req *http.Request) {
	    w.Write([]byte("Hello"))
    }

    func main() {
	    http.HandleFunc("/hello", SayHello)
	    http.ListenAndServe(":8001", nil)

    }

在d:/mnt目录下新建main.go，然后运行

    docker run --rm -v /mnt-linux:/usr/src/myapp -w /usr/src/myapp -e GOOS=windows -e GOARCH=386 golang:1.7rc1 go build -v

linux版本

    docker run --rm -v /mnt-linux:/usr/src/myapp -w /usr/src/myapp golang:1.7rc1 go build -v

在浏览器输入：

http://127.0.0.1:8001/hello

boot2docker缺省的用户名是docker，密码是tcuser


## 下载开发及编译环境

docker pull daocloud.io/library/golang:1.7-onbuild



    docker run --rm -v /mnt-linux/tx:/usr/src/myapp -w /usr/src/myapp daocloud.io/library/golang:1.7-onbuild go build -v



原因是我们的main文件生成的时候依赖的一些库如libc还是动态链接的，但是scratch 镜像完全是空的，什么东西也不包含，所以生成main时候要按照下面的方式生成，使生成的main静态链接所有的库：

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .



### 依赖管理

维护manifest依赖：

    gvt restore -connections 8

目录结构如下:
linux:
 -GOPATH:
 --src
    --vendor
      --manifest


提交已经下载好依赖的镜像

    docker commit -m='gvt-anooc' --author='zouhuigang' 614122c0aabb gvt-anooc

基于依赖，运行程序

    docker run -i -t -v /mnt-linux/anooc_go/app:/go/src/app gvt-anooc /bin/bash