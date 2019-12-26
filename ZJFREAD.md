#前言
frp 是一个高性能的反向代理应用，可以轻松地进行内网穿透，对外网提供服务，支持 TCP、UDP、HTTP、HTTPS 等协议类型，并且 web 服务支持根据域名进行路由转发。
Github: https://github.com/fatedier/frp
当然frp作者已经提供多达20种已编译好的各种版本可以供大家使用，几乎不需要自己编译。
但是有时候我们需要自定义一些内容，这时候就需要自行编译了。
搭建GO环境

##1、下载二进制包：go1.11.5.linux-amd64.tar.gz。
cd /root/
wget  https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz

##2、将下载的二进制包解压至 /usr/local目录。
tar -C /usr/local -xzf  go1.11.5.linux-amd64.tar.gz

##3、将 /usr/local/go/bin 目录添加至PATH环境变量：
vim  /etc/profile

##最后增加如下一行
export PATH=$PATH:/usr/local/go/bin

##直接运行以下命令
source /etc/profile
go version
cd  /root/
go get github.com/scientistzjf/frp
cd /root/go/src/github.com/scientistzjf/frp
make -f Makefile.cross-compiles

##~~##修改404页面~~
~~404页面在源码路径utils/vhost/resource.go文件里~~
~~改完之后再从新编译下就可以了。~~

##404html模板
```
<!DOCTYPE HTML>  
<html>  
<head>  
<meta charset="UTF-8" />  
<meta name="viewport" content="width=device-width, initial-scale=1">  
<meta name="robots" content="none" />  
<title>404 Not Found</title>  
<style>  
*{font-family:"Microsoft Yahei";margin:0;font-weight:lighter;text-decoration:none;text-align:center;line-height:2.2em;}  
html,body{height:100%;}  
h1{font-size:100px;line-height:1em;}  
table{width:100%;height:100%;border:0;}  
</style>  
</head>  
<body>  
<table cellspacing="0" cellpadding="0">  
<tr>  
<td>  
<table cellspacing="0" cellpadding="0">  
<tr>  
<td>  
<h1>404</h1>  
<h3>大事不妙啦！</h3>  
<p>你访问的页面好像不小心被博主给弄丢了~<br/>  
<a href="/">惩罚博主 ></a>  
</p>  
</td>  
</tr>  
</table>  
</td>  
</tr>  
</table>  
</body>  
</html>  
```
