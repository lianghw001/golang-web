# golang-web

## 开发 web 服务程序

开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

#### 基本要求
- 编程 web 服务程序 类似 cloudgo 应用。
- 使用 curl 测试，将测试结果写入 README.md
- 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。
- 扩展：分析net/http 源码，解释一些关键功能实现

#### 实验结果：

- 运行web服务[参考博客](https://blog.csdn.net/pmlpml/article/details/78404838)

![运行web服务器](https://github.com/lianghw001/golang-web/blob/master/pictures/service.PNG)



- curl 测试[参考博客](https://blog.csdn.net/iamlihongwei/article/details/73743278)

`$curl -v http://localhost:9090/Hello-web/whale`

![curl 测试](https://github.com/lianghw001/golang-web/blob/master/pictures/curl.PNG)




- ab 测试[参考博客](https://www.jianshu.com/p/3fd8ac3b937c)

`$ ab -n 1000 -c 100 http://localhost:9090/hello/your`

![ab 测试](https://github.com/lianghw001/golang-web/blob/master/pictures/abTest.PNG)

  - ab测试分析
    - flag“-n” 全部请求数
    - flag“-c” 并发数
    
````
Concurrency Level:      100           //并发数  
Time taken for tests:   10.68 seconds //全部请求完成耗时
Complete requests:      1000         //全部请求数
Failed requests:        0           //失败的请求
Total transferred:      176000bytes  //总传输大小 
HTML transferred:       53000bytes //整个场景中的HTML内容传输量
Requests per second:    936.76 [#/sec] (mean)   //每秒请求数(平均) 
Time per request:       106.751 [ms] (mean)   //每次并发请求时间
Transfer rate:         161.01 [Kbytes/sec] received    //传输速率
Percentage of the requests served within a certain time (ms)
  50%   107
  66%   108
  75%   109
  80%  109
  90%  111
  95%  112
  98%  115
  99%  122
 100%  210 (longest request)
//整个场景中所有请求的响应情况。在场景中每个请求都有一个响应时间，其中50％的用户响应时间小于107毫秒，60％ 的用户响应时间小于108毫秒，最大的响应时间小于210毫秒
````


- [分析net/http 源码](https://blog.csdn.net/lhw1233333/article/details/84145964)





