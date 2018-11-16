# golang-web

## 开发 web 服务程序

开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

#### 基本要求
- 编程 web 服务程序 类似 cloudgo 应用。
- 使用 curl 测试，将测试结果写入 README.md
- 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。
- 扩展：分析net/http 源码，解释一些关键功能实现

#### 实验结果：

- 运行web服务

[参考博客](https://blog.csdn.net/pmlpml/article/details/78404838)

- curl 测试
[参考博客](https://blog.csdn.net/iamlihongwei/article/details/73743278)
curl -v http://localhost:9090/Hello-web/whale

- ab 测试

ab -n 1000 -c 100 http://localhost:9090/hello/your
[参考博客](https://www.jianshu.com/p/3fd8ac3b937c)


- [分析net/http 源码](https://github.com/lianghw001/golang-web/edit/master/README.md)





