# BiliBili助手全设备支持 
Centos，Ubuntu等Linux系统,苹果Mac电脑,Openwrt路由器,Windows32位,Windows64位（一般用这个）

# [主项目地址](https://github.com/tobycroft/BiliHP-APP)

# 项目运行周期

    项目开始时间：2020-1-18  12:22:11
    项目预计结束时间：2028-03-30 12:22:11
    
## **重启原因**

学了Golang，要找个项目练练手

## **项目介绍**
助手PC版本，这是一个全新的概念版，本版本将会参考BaiduPCS-GO的web版本

不带界面，但是开放第三方API

2017年，我在客户端做好所有的界面和接口，大家直接用就好了

2020年，我打算在做好界面的基础上，向大家开放DIY接口，你可以直接用本项目的html文件实现DIY，
甚至，如果你愿意，你可以在公网上开放这个接口，或者制作一个集群管理工具，因为都是基于HTTP的，
所以技术难度很低

原来我是我带大家玩，现在我希望大家可以一起来玩，欢迎在本版本开发完成后提交branch！

![app](https://github.com/tobycroft/BiliHP-APP/blob/master/res/github-app.png)

## 版本号说明
### 例如v1.12.33：
##### 首先将版本号切为三段：
- 第一段v1为大版本号，一般是架构变化的时候才会+1，测试版V0，正式版V1，后期架构变化前段+1

- 第二段12为功能版本号，例如新增了功能，中段版本号+1

- 第三段33为修复版本号，例如12版本中出现了BUG，后期修复后发布新版本时，后段+1

- 后段无论升级到什么版本，只要前一段+1，例如功能更新，则默认新功能的包中包含之前已经修复好的bug，
如果上一个版本中仍旧有BUG，则下个版本中修复后，再下个版本的三段号中+1，以此类推


## 编译版本说明
Centos，Ubuntu等Linux系统	：BiliHP_Linux_linux

苹果Mac电脑					：BiliHP_Mac_darwin

Openwrt路由器				：BiliHP_Router_linux

Windows32位					：BiliHP_PCWEB_386

Windows64位（一般用这个）	：BiliHP_PCWEB



Windows版本和linux等新版本均不会弹出浏览器（因为新增路由器版本），请手动访问

http://127.0.0.1

或者

http://localhost


如果部署在路由器，请访问：

http://127.0.0.1:79

### DevLog
~~~~
v0.22.0
1.天选屏蔽词
~~~~
~~~~
v0.20.3
1.拆包粘包算法按照C#版本的修复
~~~~
~~~~
v0.20.2
1.本地化验证码，避免部分用户验证码死活刷不出来
~~~~
~~~~
v0.20.0
1.升级BiliHP全系列的登陆方案，解决异地登陆可能造成的白嫖风险
~~~~
~~~~
v0.20.0
1.升级BiliHP全系列的登陆方案，解决异地登陆可能造成的白嫖风险
~~~~
~~~~
v0.19.0
1.新增远程监控，从APP上就能看到PC的运行情况了，远程设定，远程查看全部完成！
（非常方便！而且如果你同时使用C2C和PC，那么监控面板将会同时显示C2C和PC的内容）
~~~~
~~~~
v0.18.11
1.修正断线动作
~~~~
~~~~
v0.18.9
1.修复SuperCurl导致故障的BUG
~~~~
~~~~
v0.18.8
1.重写拆包粘包算法，真尴尬，当初写的时候不仔细，今天花了2个小时对数据包进行挨个分析才找到拆包故障所在
大家在写拆包算法的时候要注意末尾补位问题，另外TCP粘包算法比较简单，UDP
~~~~
~~~~
v0.18.7
1.放弃自适应MSS，老老实实采用粘包方案，好气啊！
~~~~
~~~~
v0.18.1-v0.18.5
1.修复节奏风暴闪电网络版
~~~~
~~~~
v0.18.1-v0.18.5
1.干了傻事要弥补（服务器地址写成了127.0.0.1）
~~~~
~~~~
v0.18.0
1.新增节奏风暴闪电网络版
~~~~
~~~~
v0.17.2
1.修正MSS算法
~~~~
~~~~
v0.16.0-v0.17.0
1.修复关闭小电视PK舰长等所有动态奖励后仍旧领取的bug
2.新增PK/大乐斗抽奖
3.新增节奏风暴（等后端开启后下个版本更新按钮）
4.新增提督/拿督抽奖（总督，舰长，拿督，提督，目前已经做全了，之前只能抽舰长和总督）
5.新增联动开关更新

服务器：
1.修复PK模板
2.修复舰长模板
3.PC/APP接口模拟方案区分，发送方案全模拟（Cookie深度模拟将会在研究后再启用，目前GayHub独一份）
4.自动更新系统上线
5.节奏风暴模板编辑中
6.C2CGo与闪电网络对接完成，目前开放各分区前150个直播间，主网上线后将会根据C2C在线数量动态调节

~~~~
~~~~
v0.15.3
1.修复闪电网络崩坏时出现的报错退出BUG
~~~~
~~~~
v0.15.2
1.修复忘记复制html的问题
~~~~
~~~~
v0.15.1
1.修复版本忘记填写的问题
~~~~
~~~~
v0.15.0
1.新增天选之子
2.新增linux版本mac版本和openwrt路由器版本
~~~~
~~~~
v0.14.0
1.上一个版本太强大，导致程序疯狂进入退出，短时间内创造了几万个TCP链接，服务器就出现了Goroutine Map
抢资源的问题，然后通过学习，解决了Map多线程问题，老版本不会再影响服务器，导致服务器炸裂了，然后14.0版本
加入了版本，预计下个版本直接加入自动更新
~~~~
~~~~
v0.13.6
1.修复因为没有设定GoRoutine导致的协程阻塞问题
~~~~
~~~~
v0.13.5
1.html部分微调
~~~~

~~~~
v0.13.4
1.修复重连BUG
~~~~

~~~~
v0.13.1-v0.13.3
1.新增32位版本（release
2.修复登录BUG
~~~~
~~~~
v0.13.0  
    1.目前完成了礼物接入功能
TODO//
目前需要完成基础功能接入，因为群内投票，大家还是喜欢控制本地化，所以这里需要在程序中集成发起
不过目前界面登录这些可以用了，就先发了
版本同APP
~~~~
~~~~
2020-2-4:目前未完成
~~~~
