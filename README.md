# ali_update

    区别于部分 DDNS 服务，大部分供应商给出的为单一 DoMain 下基本只允许只有一个公网 IP 被长效连接。
    此项目主要供给于只有一个域名，但有公网动态 IP 的情况下，实时获取自己公网 IP 并实时更新解析道域名服务器上。

## 功能

    预计提供两类功能：
    
    1。 单点式上报更新。区分为 配置类型 和 网络读取类型。
    2。 C-S 类型。统一由 Server 上报更新，由 Client 单点触发。

## 依赖

    需要路由口有公网 IP！！！
    需要路由口有公网 IP！！！
    需要路由口有公网 IP！！！
    此项目目前仅接入 万网 提供的 SDK。

## 声明

    此项目不做商用用途！！！！