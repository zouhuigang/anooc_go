

##gvt包管理使用方法

###下载 gvt 依赖管理工具

    go get github.com/polaris1119/gvt

###进入项目目录src/
   
    gvt fetch 包路径

例如：
 
    gvt fetch github.com/fatih/color

getpkg脚本是为了一键下载manifest中定义的包


##动态设置go的运行目录update-gopath.bat

以管理员权限运行


=====此项目已废弃,仅供初学者参考，原因是已移动到docker中开发，目录变动较大，新地址后续发布=====

如您要使用：请先进入/config/env.ini修改数据库及密码或者删除env.ini，重新安装项目