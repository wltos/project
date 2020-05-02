# 一、简介
最近，我无意中看到了一个好玩的项目，即“Server酱”。于是，我把它跟爬虫技术结合到了一起，然后就有了本项目——微博热搜自动推送程序。

# 二、开发心得
1. 准备好github账号，用它换取SCKEY。有了SCKEY方可使用Server酱的推送服务；
2. 准备好微信，扫码绑定；
3. Golang知识，因为本程序是golang写成。理论上你只需换一个SCKEY即可上线；
4. HTML基础知识；
5. Server酱的官网的A&Q里指出，它并不是无限使用的。我初略算了下，一天最多请求1000次，也就是最多2分钟请求一次，如果大于这个频率，会被Server酱拉黑，解除的途径是向作者支付50块罚款或者换账号。对于这一点，我建议用redis加持一下，统计推送次数，隔夜清零。

# 三、采用的技术栈
- [ServerChan](http://sc.ftqq.com/3.version)
- Golang

# 四、架构图
![架构图](https://github.com/wltos/project/blob/feature/wei-bo-top/assets/20200502_01.jpg?raw=true)

# 五、运行程序
```
// 下载
> git clone https://github.com/wltos/project.git
> cd project
> git checkout feature/wei-bo-top

// 修改
> vi /configure/default.toml
Key="改成你自己的SCKEY"

// 运行
> go run main.go
```
# 六、效果图
## 6.1 推送
![推送](https://github.com/wltos/project/blob/feature/wei-bo-top/assets/20200502_02.jpg?raw=true)

## 6.2 推送详情
![推送详情](https://github.com/wltos/project/blob/feature/wei-bo-top/assets/20200502_03.jpg?raw=true)
