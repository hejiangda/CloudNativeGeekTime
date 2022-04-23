# 第八周作业第二部分

除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。
来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧。
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）：

Service
Ingress
可以考虑的细节

如何确保整个应用的高可用。
如何通过证书保证 httpServer 的通讯安全。

1. Service
答：
    见httpserver-svc.yaml，写一个kind为Service的yaml，type可为
    - ClusterIP
    - NodePort
    - LoadBalancer
    - ExternalName
2. Ingress
答：
   - helm安装：`curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash`
   - 没有EIP的话可以修改：`kubectl edit svc ingress-nginx-controller -n ingress`
     添加 
   ```yaml
   externalIPs: 
   - 192.168.0.10
   ```
   - 证书 
   用的之前在阿里云买的域名，没备案，采用dns方式签发证书
3. 保证应用的高可用
答：
   创建多个pod