# 第三周作业

1. 构建本地镜像
   
   答：
   1. 写好Dockerfile
   2. sudo docker build .

2. 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化

    答：Dockerfile文件在homework文件夹下

3. 将镜像推送至 docker 官方镜像仓库
    
   答：
   1. 注册一个docker hub 的帐号
   2. sudo docker login 登录
   3. 对已经创建好的镜像改名，于我而言改成 hejiangda/cloud-native-geektime-camp:xxxx
   4. sudo docker push hejiangda/cloud-native-geektime-camp:week3-homework
    
4. 通过 docker 命令本地启动 httpserver
   
   答：
   1. sudo docker run -p 8080:8080 hejiangda/cloud-native-geektime-camp:week3-homework
   2. curl localhost:8080/healthz 返回 working

   3. 通过 nsenter 进入容器查看 IP 配置
   
      答：
      1. sudo lsns | grep httpserverHomework 找到对应容器进程的pid为 106584
      2. sudo nsenter -t 106584 -n ip addr
      3. 返回结果如下：
         ```
         hjd@hasee:~/CloudNativeGeekTime$ sudo nsenter -t 106584 -n ip addr
         1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
             link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
             inet 127.0.0.1/8 scope host lo
                valid_lft forever preferred_lft forever
         55: eth0@if56: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
             link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
             inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
                valid_lft forever preferred_lft forever
         ```

作业需编写并提交 Dockerfile 及源代码。