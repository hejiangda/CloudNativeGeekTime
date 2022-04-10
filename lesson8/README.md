# 第八周作业

第一部分
   现在你对 Kubernetes 的控制面板的工作机制是否有了深入的了解呢？
   是否对如何构建一个优雅的云上应用有了深刻的认识，那么接下来用最近学过的知识把你之前编写的 http 以优雅的方式部署起来吧，你可能需要审视之前代码是否能满足优雅上云的需求。
   作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。

1. 优雅启动
答：
   1. readinessProbe 判断是否已经就绪
   2. startupProbe 慢启动
2. 优雅终止
答：
   1. 对第二周的httpserver进行了更新，实现了对SIGTERM事件的监听
   2. perStop 可以执行终止的脚本，这里因为我实现了SIGTERM事件的监听，就没启用
   3. terminationGracePeriodSeconds: 30 30秒后还没终止就开杀
3. 资源需求和 QoS 保证
答：
   1. resources 设置相关cpu和内存的限制
4. 探活
答：
   1. livenessProbe 探活
5. 日常运维需求，日志等级
答：
   1. revisionHistoryLimit: 10 限制记录的pod版本数
6. 配置和代码分离
答：
   1. ConfigMap myenv.yaml设置httpport环境变量，容器中的httpserver获取环境变量后监听相应端口启动服务
