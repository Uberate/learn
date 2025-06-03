+++
date = '2025-06-03T23:13:49+08:00'
title = '容器网络 - IpTables'
tags = ["网络", "容器", "IpTables"]
categories = ["技术"]
+++

# IpTables - Linux 中的强力防火墙

`IpTables` 是一个 `NetFilter` 的用户态应用程序。

主要面向系统管理员的 Linux 内核防火墙。

> 是的，`IpTables` 实际是一个防火墙，而不是路由器。

## 容器时代的机遇

> 订书器也没有想到有一天他订装的不是精神粮食，而是真的粮食。

除网络隔离之外，IpTables 还提供了流量隔离、管理的能力。

## IpTables 的 Table(表)

容器时代更加常见使用的是 `NAT` 表，比如 `Docker` 将服务暴露到主机服务上，实际上就是通过 `IpTables` 实现的。
如果在启动 `Docker` 添加启动参数 `--iptables=false`, `Docker` 就无法通过 `IpTables` 将服务暴露在主机上。

当然，你还可以通过其他工具实现这个功能。

### Filter 表

主要用于数据包的过滤：
- 允许或拒绝特定IP、端口的访问。
- 限制流量速率
- 基于链接状态（Established、Related）放行合法流量

```bash
# 允许已建立的连接继续通信
iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
# 拒绝所有外部主机访问本机的TCP 22端口（SSH）
iptables -A INPUT -p tcp --dport 22 -j DROP
```

### NAT 表
实现网络地址转换（NAT），修改数据包的源 / 目的 IP 或端口：
- 端口映射（如将宿主机 8080 端口映射到容器的 80 端口）。
- SNAT（源 NAT，让内网主机通过公网 IP 访问外部）。
- DNAT（目的 NAT，将外部流量转发到内网特定主机）。

```bash
# 端口映射：将外部访问宿主机8080端口的流量转发到内部服务器的80端口
iptables -t nat -A PREROUTING -p tcp --dport 8080 -j DNAT --to-destination 192.168.1.100:80
# SNAT：让内网192.168.1.0/24网段通过宿主机公网IP访问外部
iptables -t nat -A POSTROUTING -s 192.168.1.0/24 -o eth0 -j MASQUERADE
```
### Mangle 表

修改数据包的 IP 头信息（如 TTL、TOS、MARK），用于 QoS 或流量标记:
- 修改数据包的 TTL 值（如防止 IP 被追踪）。
- 为特定流量打标记，以便后续策略路由。

```bash
# 将HTTP流量的TTL设置为64
iptables -t mangle -A PREROUTING -p tcp --dport 80 -j TTL --ttl-set 64
# 为来自192.168.1.100的数据包打标记100
iptables -t mangle -A PREROUTING -s 192.168.1.100 -j MARK --set-mark 100
```
### Raw 表

配置数据包追踪（connection tracking）的例外，用于性能优化:
- 对高流量服务（如 CDN）禁用连接追踪，减少内核开销。

```bash
# 不对ICMP流量进行连接追踪（减少资源消耗）
iptables -t raw -A PREROUTING -p icmp -j NOTRACK
iptables -t raw -A OUTPUT -p icmp -j NOTRACK
```

### Security 表

基于 SELinux（Security-Enhanced Linux）实现强制访问控制（MAC）：
- 限制容器间的通信权限（如 Kubernetes NetworkPolicy 底层可能使用）。

```bash
# 限制特定SELinux上下文的进程访问网络
iptables -t security -A OUTPUT -m selinux --selinux-user unconfined_u -j DROP
```

# 引用参考
- [WIKI-IpTables](https://en.wikipedia.org/wiki/Iptables)