# gocherry-api-gateway

Golang开发的微服务网关。实现http的接口转发，下游服务配置，接口授权校验，黑白名单校验，
服务降级开关，接口缓存，默认返回值配置，插件化开发启用禁用。友好的可配置界面，方便将接口快速配置使用。


#### 基础
golang iris redis etcd mysql

#### 为什么用gocherry
1.友好的后台配置管理  
2.后台用户权限管理  
3.应用隔离，集群配置，服务节点，接口配置  
4.接口聚合  
5.接口单转发  
6.请求类型限制  
7.接口开启禁用  
8.降级处理  
9.登录授权  
10.接口缓存   
11.单接口限流  
12.ip黑名单  
13.白名单  
14.服务发现     
15.健康检查  
15...  

#### 快速开始

1.  golang安装
2.  git clone 代码  && cd gocherry-api-gateway
3.  新建mysql数据库，并导入docs目录下的sql文件  
4.  启动etcd 启动redis  
5.  cp config/app.example.tml config/app.tml 
6.  修改config/app.tml配置 (admin proxy common 三个选项都要配置)
7.  export GO111MODULE=auto && go mod vendor 下载依赖
8.  rizla admin/main.go 热启动管理后台服务
9.  cd web/ui && npm install && npm run dev 启动管理后台前端静态页面
10.  rizla  proxy/main.go 启动proxy代理层的api转发服务
11. 访问静态页面即可进入后台页面（默认端口8081），输入默认的账号密码即可

环境  
(rizla是热启动go项目 可以直接go build main.go也可以。后续加入supervisor)    
  主要关注admin proxy web/ui 三个目录 

  ![](https://github.com/lovesgg/gocherry-api-gateway/blob/master/docs/about.png)  

#### 备注
欢迎一起探讨学习

#### 打赏

请我喝咖啡 微信: jdk010
  ![](https://github.com/lovesgg/gocherry-api-gateway/blob/master/docs/me.jpeg)  

#### 未来计划
1.增加更多插件，优化现有插件  
2.服务发现，健康检查  
3.接口聚合  
4.接口默认返回值  
5....  


