# XrayR

[![](https://img.shields.io/badge/TgChat-@XrayR讨论-blue.svg)](https://t.me/XrayR_project)
[![](https://img.shields.io/badge/Channel-@XrayR通知-blue.svg)](https://t.me/XrayR_channel)
![](https://img.shields.io/github/stars/Aqr-K/XrayR)
![](https://img.shields.io/github/forks/Aqr-K/XrayR)
![](https://github.com/Aqr-K/XrayR/actions/workflows/release.yml/badge.svg)
![](https://github.com/Aqr-K/XrayR/actions/workflows/docker.yml/badge.svg)
[![Github All Releases](https://img.shields.io/github/downloads/Aqr-K/XrayR/total.svg)]()

[English](https://github.com/Aqr-K/XrayR/blob/master/README-en.md)|[Iranian](https://github.com/Aqr-K/XrayR/blob/master/README_Fa.md)|[Vietnamese](https://github.com/Aqr-K/XrayR/blob/master/README-vi.md)

A Xray backend framework that can easily support many panels.

一个基于 Xray 的后端框架，支持 V2ay,Trojan,Shadowsocks 协议，极易扩展，支持多面板对接。

如果您喜欢本项目，可以右上角点个 star+watch，持续关注本项目的进展。

使用教程：[详细使用教程](https://Aqr-K.github.io/XrayR-doc/)

## 免责声明

本项目只是本人个人学习开发并维护，本人不保证任何可用性，也不对使用本软件造成的任何后果负责。

## 特点

- 永久开源且免费。
- 支持 V2ray，Trojan， Shadowsocks 多种协议。
- 支持 Vless 和 XTLS 等新特性。
- 支持单实例对接多面板、多节点，无需重复启动。
- 支持限制在线 IP
- 支持节点端口级别、用户级别限速。
- 配置简单明了。
- 修改配置自动重启实例。
- 方便编译和升级，可以快速更新核心版本， 支持 Xray-core 新特性。

## 功能介绍

| 功能              | v2ray | trojan | shadowsocks |
| ----------------- | ----- | ------ | ----------- |
| 获取节点信息      | √     | √      | √           |
| 获取用户信息      | √     | √      | √           |
| 用户流量统计      | √     | √      | √           |
| 服务器信息上报    | √     | √      | √           |
| 自动申请 tls 证书 | √     | √      | √           |
| 自动续签 tls 证书 | √     | √      | √           |
| 在线人数统计      | √     | √      | √           |
| 在线用户限制      | √     | √      | √           |
| 审计规则          | √     | √      | √           |
| 节点端口限速      | √     | √      | √           |
| 按照用户限速      | √     | √      | √           |
| 自定义 DNS        | √     | √      | √           |

## 支持前端

| 前端                                                        | v2ray | trojan | shadowsocks                     |
| ----------------------------------------------------------- | ----- | ------ | ------------------------------- |
| sspanel-uim                                                 | √     | √      | √ (单端口多用户和 V2ray-Plugin) |
| v2board                                                     | √     | √      | √                               |
| [PMPanel](https://github.com/ByteInternetHK/PMPanel)        | √     | √      | √                               |
| [ProxyPanel](https://github.com/ProxyPanel/ProxyPanel)      | √     | √      | √                               |
| [WHMCS (V2RaySocks)](https://v2raysocks.doxtex.com/)        | √     | √      | √                               |
| [GoV2Panel](https://github.com/pingProMax/gov2panel)        | √     | √      | √                               |
| [BunPanel](https://github.com/pennyMorant/bunpanel-release) | √     | √      | √                               |

## 软件安装

### 一键安装

```
wget -N https://raw.githubusercontent.com/Aqr-K/XrayR-release/master/install.sh && bash install.sh
```

### 使用 Docker 部署软件

[Docker 部署教程](https://Aqr-K.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/docker)

### 手动安装

[手动安装教程](https://Aqr-K.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/manual)

## 配置文件及详细使用教程

[详细使用教程](https://Aqr-K.github.io/XrayR-doc/)

## Thanks

- [Project X](https://github.com/XTLS/)
- [V2Fly](https://github.com/v2fly)
- [VNet-V2ray](https://github.com/ProxyPanel/VNet-V2ray)
- [Air-Universe](https://github.com/crossfw/Air-Universe)

## Licence

[Mozilla Public License Version 2.0](https://github.com/Aqr-K/XrayR/blob/master/LICENSE)

## Telgram

[XrayR 后端讨论](https://t.me/XrayR_project)

[XrayR 通知](https://t.me/XrayR_channel)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/Aqr-K/XrayR.svg)](https://starchart.cc/Aqr-K/XrayR)
