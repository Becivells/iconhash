# Fofa Shodan icon hash 计算器

[![Latest release](https://img.shields.io/github/v/release/becivells/iconhash)](https://github.com/becivells/iconhash/releases/latest)
[![dev build status](https://img.shields.io/travis/becivells/iconhash/dev.svg?label=travis%20dev%20build)](https://travis-ci.org/becivells/iconhash)
[![master build status](https://img.shields.io/travis/becivells/iconhash/master.svg?label=travis%20master%20build)](https://travis-ci.org/becivells/iconhash)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/becivells/iconhash)
![GitHub All Releases](https://img.shields.io/github/downloads/becivells/iconhash/total)
![GitHub issues](https://img.shields.io/github/issues/becivells/iconhash)
[![Docker Pulls](https://img.shields.io/docker/pulls/becivells/iconhash)](https://hub.docker.com/r/becivells/iconhash)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/becivells/iconhash)


## 说明

大致说一下思路首先获得 favicon.ico 文件然后进行 base64 编码，编码后的数据要求每 76 个字符加上换行符。具体原因 RFC 822 文档上有说明。然后 32 位 mmh3 hash

这里贴出来 python 的实现

Python2

```shell
alias ico_hash2="python2 -c \"import requests,sys,mmh3;print mmh3.hash(requests.get(sys.argv[1]).content.encode('base64'))\""
```

Python3

```shell
alias ico_hash=" python3 -c \"import requests,sys,mmh3,codecs;print('icon_hash=%s'%mmh3.hash(codecs.lookup('base64').encode(requests.get(sys.argv[1]).content)[0]))\""
```



## 安装下载

[跳转到下载链接](https://github.com/becivells/iconhash/releases/latest)

下载后记得校验文件

```
sha1sum -c *.shasum
```



## iconhash 用法

### 不带参数

#### 1. 不带参数的 URL

```shell
iconhash https://www.baidu.com/favicon.ico
```

结果：

```
-1507567067
```

#### 2. 不带参数的 File

```shell
iconhash favicon.ico
```

**结果：**

```
-1507567067
```

**需要注意不指明文件类型 iconhash 所有的参数都不能使用**

### 带参数

#### 1. Icon 文件的 hash 值

```shell
iconhash -file favicon.ico
```

**结果：**

```
icon_hash="-1507567067"
```

注意默认使用 Fofa 的搜索格式如果关闭可以使用

```shell
iconhash -file favicon.ico -fofa=false
```

**结果：**

```
-1507567067
```



如果**只需要** shodan 格式

```shell
iconhash -file favicon.ico -fofa=false -shodan
```

**结果：**

```
http.favicon.hash:-1507567067
```

#### 2. 链接中的 icon hash 值

```
iconhash -url https://www.baidu.com/favicon.ico
```

#### 3. base64 的 icon hash 值

格式如下所示

```
data:image/vnd.microsoft.icon;base64,AAABAAEAQEAAAAEAIAAoQgAAFgAAACgAAABAAAAAgAAAAAEAIAAAAAAAAEAAAAAA
.....
//////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=
```

**执行命令**

```
iconhash -b64 imgb64.txt
```

**结果：**

```
-1507567067
```

对于这种类型的数据如果计算不准确可以去掉`data:image/vnd.microsoft.icon;base64,`只保留 base64 数据试试

#### uint32数据

默认使用的是 int32 数据如果想获得 uint32 的值可以加参数 -uint32



## 帮助

```
Usage of ./iconhash:
  -b64 string
        mmh3 hash image base64 from file 
         iconhash   -file test/favicon.ico 
  -file string
        mmh3 hash from file 
         iconhash -file favicon.ico
  -fofa
        fofa search format (default true)
  -h    look help 
         iconhash  favicon.ico 
         iconhash  https://www.baidu.com/favicon.ico
  -shodan
        shodan search format 
         iconhash   -file test/favicon.ico -shodan -fofa=false
  -uint32
        uint32
  -url string
        mmh3 hash from url 
         iconhash -url  https://www.baidu.com/favicon.ico
  -user-agent string
        mmh3 hash from url (default "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11")
  -v    version
```



### 调试模式

主要应对 url 和 image base64 的情况，如果 hash 值不一致请开启 debug 模式看错误信息

```
iconhash -url https://106.55.12.93/favicon.ico1  -debug
```

**结果：**

```
---------------------------     var    value     --------------------------------
h                  :false
v                  :false
Version            :2020-05-25 12:02:03 +0800 v0.2-11-gfcbf179
VERSION_TAG        :v0.2-11-gfcbf179
Compile            :2020-05-25 21:51:38 +0800
Branch             :dev
GitDirty           :73
HashUrl            :https://106.55.12.93/favicon.ico1
Hashfile           :
ImageBase64        :
UserAgent          :Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11
IsUint32           :false
FofaFormat        :true
ShodanFormat       :false
InsecureSkipVerify :true
Debug              :true
DefaultUA          :Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11
---------------------------     var    value     --------------------------------
---------------------------  start url  content  --------------------------------
====> url: https://106.55.12.93/favicon.ico1
===> status code: 404
====> content: 
<html>
<head><title>404 Not Found</title></head>
<body>
<center><h1>404 Not Found</h1></center>
<hr><center>nginx</center>
</body>
</html>
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->

---------------------------  end url  content  --------------------------------
---------------------------start base64 content--------------------------------
====> base64:
PGh0bWw+DQo8aGVhZD48dGl0bGU+NDA0IE5vdCBGb3VuZDwvdGl0bGU+PC9oZWFkPg0KPGJvZHk+
DQo8Y2VudGVyPjxoMT40MDQgTm90IEZvdW5kPC9oMT48L2NlbnRlcj4NCjxocj48Y2VudGVyPm5n
aW54PC9jZW50ZXI+DQo8L2JvZHk+DQo8L2h0bWw+DQo8IS0tIGEgcGFkZGluZyB0byBkaXNhYmxl
IE1TSUUgYW5kIENocm9tZSBmcmllbmRseSBlcnJvciBwYWdlIC0tPg0KPCEtLSBhIHBhZGRpbmcg
dG8gZGlzYWJsZSBNU0lFIGFuZCBDaHJvbWUgZnJpZW5kbHkgZXJyb3IgcGFnZSAtLT4NCjwhLS0g
YSBwYWRkaW5nIHRvIGRpc2FibGUgTVNJRSBhbmQgQ2hyb21lIGZyaWVuZGx5IGVycm9yIHBhZ2Ug
LS0+DQo8IS0tIGEgcGFkZGluZyB0byBkaXNhYmxlIE1TSUUgYW5kIENocm9tZSBmcmllbmRseSBl
cnJvciBwYWdlIC0tPg0KPCEtLSBhIHBhZGRpbmcgdG8gZGlzYWJsZSBNU0lFIGFuZCBDaHJvbWUg
ZnJpZW5kbHkgZXJyb3IgcGFnZSAtLT4NCjwhLS0gYSBwYWRkaW5nIHRvIGRpc2FibGUgTVNJRSBh
bmQgQ2hyb21lIGZyaWVuZGx5IGVycm9yIHBhZ2UgLS0+DQo=

---------------------------end base64 content--------------------------------
icon_hash="566218143"
```



## 编译 or 开发

安装 make 环境，golang 环境，测试环境 [bats](https://github.com/bats-core/bats-core)

本地编译

```
make
```

生成 release 版本

```
make release #编译版本
```

bat 测试功能

```
make test-cli
```

