# XrayR

A Xray backend framework that can easily support many panels.

A back -end framework based on XRAY supports V2ay, Trojan, Shadowsocks protocols, which are easy to expand and support multi -panel docker.

If you like this project, you can click STAR+WATCH in the upper right corner to continue to pay attention to the progress of this project.

## Disclaimer

This project is just my personal learning and development and maintenance. I do not guarantee any availability and is not responsible for any consequences caused by the use of this software.

## Features

- Permanent open source and free.
- Support V2Ray, Trojan, Shadowsocks multiple protocols.
- Support new features such as Vless and XTLS.
- Support single instance docking multi -panel and multi -node, no need to start repeatedly.
- Support restriction online IP
- Support node port level and user level speed limit.
- The configuration is simple and clear.
- Modify the automatic restart instance.
- Easy to compile and upgrade, you can quickly update the core version and support the new features of XRAY-CORE.

## Function

| Function                                  | v2ray | trojan | shadowsocks |
| ----------------------------------------- | ----- | ------ | ----------- |
| Get node information                      | √     | √      | √           |
| Get user information                      | √     | √      | √           |
| User traffic statistics                   | √     | √      | √           |
| Server information report                 | √     | √      | √           |
| Automatically apply for a TLS certificate | √     | √      | √           |
| Automatic renewal TLS certificate         | √     | √      | √           |
| Number of online people                   | √     | √      | √           |
| Online user restrictions                  | √     | √      | √           |
| Audit rules                               | √     | √      | √           |
| Node port speed limit                     | √     | √      | √           |
| According to user speed limit             | √     | √      | √           |
| Custom DNS                                | √     | √      | √           |

## Support for panels

| Panel                                                       | v2ray | trojan | shadowsocks                                  |
| ----------------------------------------------------------- | ----- | ------ | -------------------------------------------- |
| sspanel-uim                                                 | √     | √      | √ (Single-ended multi-user and V2Ray-Plugin) |
| v2board                                                     | √     | √      | √                                            |
| PMPanel                                                     | √     | √      | √                                            |
| ProxyPanel                                                  | √     | √      | √                                            |
| V2RaySocks                                                  | √     | √      | √                                            |
| BunPanel                                                    | √     | √      | √                                            |

## Software Installation

### 1-Click installation

```
wget -N https://raw.githubusercontent.com/Sysrous/XrayR/refs/heads/master/install.sh && bash install.sh
```



## Thanks

- [Project X](https://github.com/XTLS/)
- [V2Fly](https://github.com/v2fly)
- [VNet-V2ray](https://github.com/ProxyPanel/VNet-V2ray)
- [Air-Universe](https://github.com/crossfw/Air-Universe)

## Licence

[Mozilla Public License Version 2.0](https://github.com/Aqr-K/XrayR/blob/master/LICENSE)

